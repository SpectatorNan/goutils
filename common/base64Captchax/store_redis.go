package base64Captchax

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	cacheKeyCaptchaIdPrefix = "cache:captcha:id:"
)

type redisStore struct {
	redis      *redis.Redis
	expiration int
	prefix     string
}

func NewRedisStore(redis *redis.Redis, expiration int, prefix string) base64Captcha.Store {
	s := new(redisStore)
	s.redis = redis
	s.expiration = expiration
	s.prefix = prefix
	return s
}

func (s *redisStore) Set(id string, value string) error {
	cacheKey := fmt.Sprintf("%s%s%v", cacheKeyCaptchaIdPrefix, s.prefix, id)
	return s.redis.Setex(cacheKey, value, s.expiration)
}
func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
func (s *redisStore) Get(id string, clear bool) (value string) {
	cacheKey := fmt.Sprintf("%s%s%v", cacheKeyCaptchaIdPrefix, s.prefix, id)
	value, err := s.redis.Get(cacheKey)
	if err != nil {
		return
	}
	if clear {
		_, _ = s.redis.Del(cacheKey)
	}
	return
}
