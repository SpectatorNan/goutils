package base64Captchax

import (
	"context"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	cacheKeyCaptchaIdPrefix = "cache:captcha:id:"
)

type redisStore struct {
	redis      *redis.Client
	expiration time.Duration
	prefix     string
}

func NewRedisStore(redis *redis.Client, expiration int, prefix string) base64Captcha.Store {
	s := new(redisStore)
	s.redis = redis
	s.expiration = time.Second * time.Duration(expiration)
	s.prefix = prefix
	return s
}

func (s *redisStore) Set(id string, value string) error {
	cacheKey := s.generateKey(id)
	return s.redis.SetEx(context.Background(), cacheKey, value, s.expiration).Err()
}
func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
func (s *redisStore) Get(id string, clear bool) (value string) {
	cacheKey := s.generateKey(id)
	value, err := s.redis.Get(context.Background(), cacheKey).Result()
	if err != nil {
		return
	}
	if clear {
		_, _ = s.redis.Del(context.Background(), cacheKey).Result()
	}
	return
}

func (s *redisStore) generateKey(id string) string {
	return fmt.Sprintf("%s%s%v", s.prefix, cacheKeyCaptchaIdPrefix, id)
}

// WithPrefix sets the prefix for the redis store
// rpc 服务共用的时候会出现竞争问题， 不建议使用
func withPrefix(s *redisStore, prefix string) base64Captcha.Store {
	s.prefix = prefix
	return s
}
