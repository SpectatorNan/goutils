package cachex

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type RedisSingleNodeConf struct {
	Host     string
	Pass     string `json:",optional"`
	Tls      bool   `json:",optional"`
	NonBlock bool   `json:",default=true"`
	// PingTimeout is the timeout for ping redis.
	PingTimeout time.Duration `json:",default=1s"`
}

func (conf RedisSingleNodeConf) ToG0RedisConf() redis.RedisConf {
	return redis.RedisConf{
		Host:        conf.Host,
		Type:        "node",
		Pass:        conf.Pass,
		Tls:         conf.Tls,
		NonBlock:    conf.NonBlock,
		PingTimeout: conf.PingTimeout,
	}
}

func (conf RedisSingleNodeConf) ToG0CacheConf() cache.CacheConf {
	return []cache.NodeConf{
		{
			RedisConf: redis.RedisConf{
				Host:        conf.Host,
				Type:        "node",
				Pass:        conf.Pass,
				Tls:         conf.Tls,
				NonBlock:    conf.NonBlock,
				PingTimeout: conf.PingTimeout,
			},
			Weight: 100,
		},
	}
}

func (conf RedisSingleNodeConf) ToGoRedis(opts ...ToGoRedisOptions) *goredis.Client {
	opt := &goredis.Options{
		Addr:     conf.Host,
		Password: conf.Pass,
		DB:       0,
	}

	for _, option := range opts {
		option(opt)
	}

	rdb := goredis.NewClient(opt)
	return rdb
}
