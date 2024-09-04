package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
)

type JSConfig struct {
	AccessKey    string
	AccessSecret string
}

type JSClient struct {
	accessKey    string
	accessSecret string
}

func NewJSClient(cfg JSConfig) *JSClient {
	return &JSClient{
		accessKey:    cfg.AccessKey,
		accessSecret: cfg.AccessSecret,
	}
}

func (c *JSClient) GetJSCredentials(effectiveTime time.Duration) (*OssJsCredential, error) {
	ossAppKey := c.accessKey
	ossSercret := c.accessSecret

	expiration := time.Now().Add(effectiveTime).Format("2006-01-02T15:04:05.000Z")

	policyMap := map[string]interface{}{
		"expiration": expiration,
		"conditions": []interface{}{
			[]interface{}{
				"content-length-range",
				0,
				1048576000,
			},
		},
	}
	jsonStr, err := jsonx.Marshal(policyMap)
	if err != nil {
		return nil, err
	}
	policy := base64.StdEncoding.EncodeToString(jsonStr)
	mac := hmac.New(sha1.New, []byte(ossSercret))
	mac.Write([]byte(policy))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return &OssJsCredential{
		Policy:         policy,
		OSSAccessKeyId: ossAppKey,
		Signature:      res,
	}, nil
}
