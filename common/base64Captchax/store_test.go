package base64Captchax

import (
	"github.com/redis/go-redis/v9"
	"testing"
)

func Test_StoreModifyPrefix(t *testing.T) {
	rds := redis.NewClient(&redis.Options{})
	captchaStore := NewRedisStore(rds, 10*60, "admin_")
	t.Logf("captchaStore prefix: %s", captchaStore.(*redisStore).prefix)
	abcStore := withPrefix(captchaStore.(*redisStore), "abc_")
	t.Logf("captchaStore prefix: %s", captchaStore.(*redisStore).prefix)
	t.Logf("abcStore prefix: %s", abcStore.(*redisStore).prefix)
}
