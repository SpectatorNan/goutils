package base64Captchax

import (
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

var (
	Expiration        = 10 * 60 // 10 minutes
	DefaultRedisStore = func(redis *redis.Client, prefix string) base64Captcha.Store {
		return NewRedisStore(redis, Expiration, prefix)
	}
)
