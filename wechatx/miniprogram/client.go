package miniprogram

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"sort"
	"strings"
)

type Client interface {
	GetAccessToken(ctx context.Context) (*string, error)
	GetJSApiTicket(ctx context.Context, token string) (*string, error)
	GetJSApiTicketCombined(ctx context.Context) (*string, error)
	JSApiTicketSignature(ticket, nonceStr, timestamp, url string) (string, error)
}

type Option func(c *client)

func WithCacheKeyPrefix(prefix string) Option {
	return func(c *client) {
		c.keyPrefix = prefix
	}
}
func WithRedisCache(rds *redis.Client) Option {
	return func(c *client) {
		c.cacheType = CacheTypeRedis
		c.redis = rds
	}
}

type client struct {
	appId     string
	secret    string
	cacheType CacheType
	memCache  map[string]CacheItem
	redis     *redis.Client
	keyPrefix string
}

func NewClient(appId, secret string, opts ...Option) Client {
	cli := newClient(appId, secret)
	for _, opt := range opts {
		opt(cli)
	}
	return cli
}
func (c *client) GetJSApiTicketCombined(ctx context.Context) (*string, error) {
	token, err := c.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	return c.GetJSApiTicket(ctx, *token)
}
func (c *client) GetAccessToken(ctx context.Context) (*string, error) {
	cacheKey := c.getKey(cacheKeyWithAccessToken)
	if cacheData, err := c.checkCache(ctx, cacheKey); err == nil && len(cacheData) > 0 {
		return &cacheData, nil
	}
	resp, err := c.getResponse(path_token, paramsAccessToken(c.appId, c.secret))
	if err != nil {
		return nil, errors.Wrapf(err, "GetAccessToken error request error")
	}
	result, err := handleResponse[AccessTokenResponse](resp)
	if err != nil {
		return nil, err
	}

	err = c.setCache(ctx, cacheKey, CacheItem{
		ExpireUnix: result.ExpiresIn - 105,
		Data:       result.AccessToken,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("GetAccessToken setCache error: %v", err)
	}
	return &result.AccessToken, nil
}

func (c *client) GetJSApiTicket(ctx context.Context, token string) (*string, error) {

	cacheKey := c.getKey(cacheKeyWithJSApiTicket)
	if cacheData, err := c.checkCache(ctx, cacheKey); err == nil && len(cacheData) > 0 {
		return &cacheData, nil
	}

	resp, err := c.getResponse(path_ticket, paramsGetTicket(token))
	if err != nil {
		return nil, errors.Wrapf(err, "GetJSApiTicket error request error")
	}
	result, err := handleResponse[JSApiTicketResponse](resp)
	if err != nil {
		return nil, err
	}
	err = c.setCache(ctx, cacheKey, CacheItem{
		ExpireUnix: result.ExpiresIn - 100,
		Data:       result.Ticket,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("GetJSApiTicket setCache error: %v", err)
	}
	return &result.Ticket, nil
}

func (c *client) JSApiTicketSignature(ticket, nonceStr, timestamp, url string) (string, error) {
	params := map[string]string{
		"jsapi_ticket": ticket,
		"noncestr":     nonceStr,
		"timestamp":    timestamp,
		"url":          url,
	}

	// Sort the parameters by key in ASCII order
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Concatenate the parameters in the format key1=value1&key2=value2
	var sb strings.Builder
	for i, k := range keys {
		if i > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(fmt.Sprintf("%s=%s", k, params[k]))
	}
	string1 := sb.String()

	// Compute the SHA1 hash of the concatenated string
	h := sha1.New()
	_, err := h.Write([]byte(string1))
	if err != nil {
		return "", err
	}
	signature := hex.EncodeToString(h.Sum(nil))

	return signature, nil
}

func newClient(appId, secret string) *client {
	return &client{
		appId:     appId,
		secret:    secret,
		cacheType: CacheTypeMemory,
		memCache:  make(map[string]CacheItem),
	}
}
