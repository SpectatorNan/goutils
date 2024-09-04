package gozero_redis

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	Expiration        = 10 * 60 // 10 minutes
	DefaultRedisStore = func(redis *redis.Redis, prefix string) base64Captcha.Store {
		return NewRedisStore(redis, Expiration, prefix)
	}
)
