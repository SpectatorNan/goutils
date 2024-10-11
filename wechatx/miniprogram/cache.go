package miniprogram

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func (c *client) getKey(key string) string {
	if len(c.keyPrefix) < 1 {
		return key
	}
	return fmt.Sprintf("%s:%s", c.keyPrefix, key)
}

func (c *client) setCache(ctx context.Context, key string, item CacheItem) error {
	switch c.cacheType {
	case CacheTypeMemory:
		c.memCache[key] = item
		return nil
	case CacheTypeRedis:
		if c.redis == nil {
			return errors.New("redis client is nil")
		}
		err := c.redis.Set(ctx, key, item.Data, time.Duration(item.ExpireUnix)*time.Second).Err()
		if err != nil {
			return errors.Wrapf(err, "setCache error")
		}
		return nil
	}
	return errors.Errorf("unsupported cache type: %v", c.cacheType)
}

func (c *client) checkCache(ctx context.Context, key string) (string, error) {
	switch c.cacheType {
	case CacheTypeMemory:
		if item, ok := c.memCache[key]; ok {
			if item.ExpireUnix < time.Now().Unix() && item.ExpireUnix != 0 {
				delete(c.memCache, key)
			} else {
				return item.Data, nil
			}
		}
		return "", nil
	case CacheTypeRedis:
		if c.redis == nil {
			return "", errors.New("redis client is nil")
		}
		resStr, err := c.redis.Get(ctx, key).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return "", nil
			}
			return "", err
		}
		return resStr, nil
		//var cacheItem CacheItem
		//perr := json.Unmarshal([]byte(resStr), &cacheItem)
		//if perr != nil {
		//
		//	return "", perr
		//}
		//return cacheItem.Data, nil

	default:
		return "", errors.Errorf("unsupported cache type: %v", c.cacheType)
	}
}
