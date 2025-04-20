package goredisLock

import (
	"context"
	_ "embed"
	"errors"
	gred "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	randomLen       = 16
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
)

var (
	//go:embed lockscript.lua
	lockLuaScript string
	lockScript    = NewScript(lockLuaScript)

	//go:embed delscript.lua
	delLuaScript string
	delScript    = NewScript(delLuaScript)
)

// A RedisLock is a redis lock.
type RedisLock struct {
	store   *gred.Client
	seconds uint32
	key     string
	id      string
}

func NewRedisLock(store *gred.Client, key string) *RedisLock {
	return &RedisLock{
		store: store,
		key:   key,
		id:    stringx.Randn(randomLen),
	}
}

func NewRedisLockWithSeconds(store *gred.Client, key string, seconds uint32) *RedisLock {
	return &RedisLock{
		store:   store,
		seconds: seconds,
		key:     key,
		id:      stringx.Randn(randomLen),
	}
}

// Acquire acquires the lock.
func (rl *RedisLock) Acquire() (bool, error) {
	return rl.AcquireCtx(context.Background())
}

// AcquireCtx acquires the lock with the given ctx.
func (rl *RedisLock) AcquireCtx(ctx context.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	resp, err := rl.ScriptRunCtx(ctx, lockScript, []string{rl.key}, []string{
		rl.id, strconv.Itoa(int(seconds)*millisPerSecond + tolerance),
	})
	if errors.Is(err, gred.Nil) {
		return false, nil
	} else if err != nil {
		logx.Errorf("Error on acquiring lock for %s, %s", rl.key, err.Error())
		return false, err
	} else if resp == nil {
		return false, nil
	}

	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	}

	logx.Errorf("Unknown reply when acquiring lock for %s: %v", rl.key, resp)
	return false, nil
}

// Release releases the lock.
func (rl *RedisLock) Release() (bool, error) {
	return rl.ReleaseCtx(context.Background())
}

// ReleaseCtx releases the lock with the given ctx.
func (rl *RedisLock) ReleaseCtx(ctx context.Context) (bool, error) {
	resp, err := rl.ScriptRunCtx(ctx, delScript, []string{rl.key}, []string{rl.id})
	if err != nil {
		return false, err
	}

	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return reply == 1, nil
}

// SetExpire sets the expiration.
func (rl *RedisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}

func (rl *RedisLock) ScriptRunCtx(ctx context.Context, script *Script, keys []string,
	args ...any) (any, error) {
	return script.Run(ctx, rl.store, keys, args...).Result()
}

// AcquireWithWait attempts to acquire the lock with blocking until success or context cancellation.
// The retryInterval parameter specifies the time to wait between retry attempts.
func (rl *RedisLock) AcquireWithWait(ctx context.Context, retryInterval time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			acquired, err := rl.AcquireCtx(ctx)
			if err != nil {
				return err
			}
			if acquired {
				return nil
			}
			// Wait for the specified interval before retrying
			timer := time.NewTimer(retryInterval)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
				// Continue the loop to try again
			}
		}
	}
}

// AcquireWithTimeout attempts to acquire the lock with blocking until success
// or until the specified timeout duration has elapsed.
func (rl *RedisLock) AcquireWithTimeout(timeout, retryInterval time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return rl.AcquireWithWait(ctx, retryInterval)
}
