package miniprogram

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestClient_GetAccessToken(t *testing.T) {
	ctx := context.Background()
	c := newClient("", "")
	AccessToken, err := c.GetAccessToken(ctx)
	if err != nil {
		t.Errorf("GetAccessToken error = %v", err)
		return
	}

	t.Logf("GetAccessToken AccessToken = %v", *AccessToken)
}

func TestClient_GetJSApiTicket(t *testing.T) {
	ctx := context.Background()
	c := newClient("", "")
	AccessToken, err := c.GetAccessToken(ctx)
	if err != nil {
		t.Errorf("GetJSApiTicket access token error = %v", err)
		return
	}
	t.Logf("GetJSApiTicket access_token = %v", *AccessToken)

	ticket, err := c.GetJSApiTicket(ctx, *AccessToken)
	if err != nil {
		t.Errorf("GetJSApiTicket ticket error = %v", err)
		return
	}

	t.Logf("GetJSApiTicket ticket = %v", *ticket)

}

func TestClient_CheckCache(t *testing.T) {

	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "yourpassword",
	})
	ctx := context.Background()

	type Case struct {
		Name       string
		Key        string
		CacheType  CacheType
		ExpireUnix int64
		Data       string
		Want       bool
	}

	cases := []Case{
		{
			Name:       "Memory",
			Key:        fmt.Sprintf("%s:%s", "forever", cacheKeyWithAccessToken),
			CacheType:  CacheTypeMemory,
			ExpireUnix: 0,
			Data:       "test",
			Want:       true,
		},
		{
			Name:       "Memory",
			Key:        fmt.Sprintf("%s:%s", "xseconds", cacheKeyWithAccessToken),
			CacheType:  CacheTypeMemory,
			ExpireUnix: time.Now().Unix() + 10,
			Data:       "test",
			Want:       true,
		},
		{
			Name:       "Memory",
			Key:        fmt.Sprintf("%s:%s", "expired", cacheKeyWithAccessToken),
			CacheType:  CacheTypeMemory,
			ExpireUnix: time.Now().Unix() - 10,
			Data:       "test",
			Want:       false,
		},
		{
			Name:       "Redis",
			Key:        fmt.Sprintf("%s:%s", "forever", cacheKeyWithAccessToken),
			CacheType:  CacheTypeRedis,
			ExpireUnix: -1,
			Data:       "test",
			Want:       true,
		},
		{
			Name:       "Redis",
			Key:        fmt.Sprintf("%s:%s", "xseconds", cacheKeyWithAccessToken),
			CacheType:  CacheTypeRedis,
			ExpireUnix: 30,
			Data:       "test",
			Want:       true,
		},
		{
			Name:       "Redis",
			Key:        fmt.Sprintf("%s:%s", "expired", cacheKeyWithAccessToken),
			CacheType:  CacheTypeRedis,
			ExpireUnix: 1,
			Data:       "test",
			Want:       false,
		},
	}
	//t.Run("load redis cache", func(t *testing.T) {
	//	c := newClient("", "")
	//	c.cacheType = CacheTypeRedis
	//	c.redis = rds
	//	for _, v := range cases {
	//		if v.CacheType == CacheTypeMemory {
	//			continue
	//		}
	//		err := c.redis.Set(ctx, v.Key, v.Data, time.Duration(v.ExpireUnix)*time.Second).Err()
	//		if err != nil {
	//			t.Errorf("load redis cache error = %v", err)
	//		}
	//	}
	//})

	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			c := newClient("", "")

			if v.CacheType == CacheTypeRedis {
				c.cacheType = CacheTypeRedis
				c.redis = rds
			} else {
				for _, v1 := range cases {
					if v1.CacheType == CacheTypeRedis {
						continue
					}
					c.memCache[v1.Key] = CacheItem{
						ExpireUnix: v1.ExpireUnix,
						Data:       v1.Data,
					}
				}
			}
			got, err := c.checkCache(ctx, v.Key)
			if err != nil {
				t.Errorf("checkCache error = %v", err)
				return
			}
			if (got != "" && !v.Want) || (got == "" && v.Want) {
				t.Errorf("checkCache = %v, want %v", got, v.Want)
			}
		})
	}
}
