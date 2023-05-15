package jwtx

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/goutils/common/cryptx"
	"github.com/SpectatorNan/goutils/common/tools"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
)

var (
	refreshCacheKeyPrefix = "cache:refreshToken:"
)

func CacheRefreshToken(ctx context.Context, redis *redis.Redis, token string, uid int64) error {
	redisKey := fmt.Sprintf("%s%s", refreshCacheKeyPrefix, token)
	return redis.SetexCtx(ctx, redisKey, fmt.Sprintf("%d", uid), 86400*7)
}

func GetUidByRefreshToken(ctx context.Context, redis *redis.Redis, token string) int64 {
	redisKey := fmt.Sprintf("%s%s", refreshCacheKeyPrefix, token)
	val, _ := redis.GetCtx(ctx, redisKey)
	uid, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		return uid
	}
	return 0
}

func DeleteRefreshToken(ctx context.Context, redis *redis.Redis, token string) error {
	redisKey := fmt.Sprintf("%s%s", refreshCacheKeyPrefix, token)
	_, err := redis.Del(redisKey)
	return err
}

func GenerateRefreshToken(uid int64) string {
	return cryptx.Md5ByString(fmt.Sprintf("%v%v", tools.Krand(14, tools.KC_RAND_KIND_ALL), uid))
}
