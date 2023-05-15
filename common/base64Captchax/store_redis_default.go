package base64Captchax

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

var (
	Expiration        = 10 * time.Minute
	DefaultRedisStore = func(redis *redis.Redis, prefix string) base64Captcha.Store {
		return NewRedisStore(redis, int(Expiration), prefix)
	}
)
