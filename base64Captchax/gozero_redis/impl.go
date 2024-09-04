package gozero_redis

import (
	base64Captchax2 "github.com/SpectatorNan/goutils/base64Captchax"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type redisStore struct {
	redis      *redis.Redis
	expiration int
	prefix     string
}

func NewRedisStore(redis *redis.Redis, expiration int, prefix string) base64Captchax2.Store {
	s := new(redisStore)
	s.redis = redis
	s.expiration = expiration
	s.prefix = prefix
	return s
}
func (s *redisStore) SetCaptcha(prefix string, id string, value string) error {
	cacheKey := base64Captchax2.GenerateKey(prefix, id)
	return s.Set(cacheKey, value)
}
func (s *redisStore) GetCaptcha(prefix string, id string, clear bool) string {
	cacheKey := base64Captchax2.GenerateKey(prefix, id)
	return s.Get(cacheKey, clear)
}
func (s *redisStore) VerifyCaptcha(prefix string, id, answer string, clear bool) bool {
	cacheKey := base64Captchax2.GenerateKey(prefix, id)
	return s.Verify(cacheKey, answer, clear)
}

func (s *redisStore) Set(id string, value string) error {
	return s.redis.Setex(id, value, s.expiration)
}
func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
func (s *redisStore) Get(id string, clear bool) (value string) {
	value, err := s.redis.Get(id)
	if err != nil {
		return
	}
	if clear {
		_, _ = s.redis.Del(id)
	}
	return
}
