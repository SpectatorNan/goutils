package cachex

import (
	"crypto/tls"
	goredis "github.com/redis/go-redis/v9"
	"time"
)

// only support go redis, not support go-zero redis
type ToGoRedisOptions func(opt *goredis.Options)
func WithRedisDB(number int) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.DB = number
	}
}
func WithRedisUsername(username string) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.Username = username
	}
}
func WithRedisPassword(password string) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.Password = password
	}
}
func WithRedisMaxRetries(maxRetries int) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.MaxRetries = maxRetries
	}
}
func WithRedisMinRetryBackoff(minRetryBackoff time.Duration) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.MinRetryBackoff = minRetryBackoff
	}
}
func WithRedisMaxRetryBackoff(maxRetryBackoff time.Duration) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.MaxRetryBackoff = maxRetryBackoff
	}
}
func WithRedisIdleConns(idleConns int) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.MinIdleConns = idleConns
	}
}
func WithRedisTlsConfig(tlsConfig *tls.Config) ToGoRedisOptions {
	return func(opt *goredis.Options) {
		opt.TLSConfig = tlsConfig
	}
}


func DefaultRedisMaxRetries() ToGoRedisOptions {
	return WithRedisMaxRetries(3)
}
func DefaultRedisIdleConns() ToGoRedisOptions {
	return WithRedisIdleConns(8)
}