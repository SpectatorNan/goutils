package base64Captchax

import (
	"fmt"
)

var (
	cacheKeyCaptchaIdPrefix = "cache:captcha:id:"
) 



func GenerateKey(prefix, id string) string {
	return fmt.Sprintf("%s%s%v", prefix, cacheKeyCaptchaIdPrefix, id)
}

// WithPrefix sets the prefix for the redis store
// rpc 服务共用的时候会出现竞争问题， 不建议使用
//func withPrefix(s *redisStore, prefix string) base64Captcha.Store {
//	s.prefix = prefix
//	return s
//}
