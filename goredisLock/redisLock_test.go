package goredisLock

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisLock(t *testing.T) {
	// Setup miniredis for testing
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	// Create redis client
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	defer client.Close()

	t.Run("NewRedisLock", func(t *testing.T) {
		lock := NewRedisLock(client, "test-key")
		assert.NotNil(t, lock)
		assert.Equal(t, "test-key", lock.key)
		assert.Equal(t, uint32(0), lock.seconds)
		assert.NotEmpty(t, lock.id)
		assert.Equal(t, randomLen, len(lock.id))
	})

	t.Run("NewRedisLockWithSeconds", func(t *testing.T) {
		lock := NewRedisLockWithSeconds(client, "test-key", 30)
		assert.NotNil(t, lock)
		assert.Equal(t, "test-key", lock.key)
		assert.Equal(t, uint32(30), lock.seconds)
		assert.NotEmpty(t, lock.id)
		assert.Equal(t, randomLen, len(lock.id))
	})

	t.Run("Acquire_Success", func(t *testing.T) {
		lock := NewRedisLockWithSeconds(client, "test-lock-1", 30)
		acquired, err := lock.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired)
	})

	t.Run("AcquireCtx_Success", func(t *testing.T) {
		lock := NewRedisLockWithSeconds(client, "test-lock-2", 30)
		ctx := context.Background()
		acquired, err := lock.AcquireCtx(ctx)
		assert.NoError(t, err)
		assert.True(t, acquired)
	})

	t.Run("Acquire_AlreadyLocked", func(t *testing.T) {
		key := "test-lock-3"
		lock1 := NewRedisLockWithSeconds(client, key, 30)
		lock2 := NewRedisLockWithSeconds(client, key, 30)

		// First lock succeeds
		acquired1, err := lock1.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired1)

		// Second lock fails
		acquired2, err := lock2.Acquire()
		assert.NoError(t, err)
		assert.False(t, acquired2)
	})

	t.Run("Release_Success", func(t *testing.T) {
		key := "test-lock-4"
		lock := NewRedisLockWithSeconds(client, key, 30)

		// Acquire the lock
		acquired, err := lock.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired)

		// Release the lock
		released, err := lock.Release()
		assert.NoError(t, err)
		assert.True(t, released)

		// Can acquire it again
		acquired2, err := lock.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired2)
	})

	t.Run("Release_WrongID", func(t *testing.T) {
		key := "test-lock-5"
		lock1 := NewRedisLockWithSeconds(client, key, 30)
		lock2 := NewRedisLockWithSeconds(client, key, 30)

		// Lock with lock1
		acquired, err := lock1.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired)

		// Try to release with lock2 (has different ID)
		released, err := lock2.Release()
		assert.NoError(t, err)
		assert.False(t, released)
	})

	t.Run("SetExpire", func(t *testing.T) {
		lock := NewRedisLockWithSeconds(client, "test-key", 30)
		assert.Equal(t, uint32(30), lock.seconds)

		lock.SetExpire(60)
		assert.Equal(t, uint32(60), lock.seconds)
	})

	t.Run("AcquireWithWait_Success", func(t *testing.T) {
		key := "test-lock-6"
		lock := NewRedisLockWithSeconds(client, key, 30)

		err := lock.AcquireWithWait(context.Background(), 10*time.Millisecond)
		assert.NoError(t, err)

		// Verify it was locked
		lock2 := NewRedisLockWithSeconds(client, key, 30)
		acquired, err := lock2.Acquire()
		assert.NoError(t, err)
		assert.False(t, acquired)
	})

	t.Run("AcquireWithWait_WaitsForRelease", func(t *testing.T) {
		key := "test-lock-7"
		lock1 := NewRedisLockWithSeconds(client, key, 30)
		lock2 := NewRedisLockWithSeconds(client, key, 30)

		// Lock with lock1
		acquired, err := lock1.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired)

		// Start goroutine to release after delay
		go func() {
			time.Sleep(50 * time.Millisecond)
			released, err := lock1.Release()
			assert.NoError(t, err)
			assert.True(t, released)
		}()

		// lock2 should wait and eventually succeed
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		err = lock2.AcquireWithWait(ctx, 10*time.Millisecond)
		assert.NoError(t, err)
	})

	t.Run("AcquireWithWait_Timeout", func(t *testing.T) {
		key := "test-lock-8"
		lock1 := NewRedisLockWithSeconds(client, key, 30)
		lock2 := NewRedisLockWithSeconds(client, key, 30)

		// Lock with lock1
		acquired, err := lock1.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired)

		// lock2 should time out
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		err = lock2.AcquireWithWait(ctx, 10*time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, context.DeadlineExceeded, err)
	})

	t.Run("AcquireWithTimeout_Success", func(t *testing.T) {
		key := "test-lock-9"
		lock := NewRedisLockWithSeconds(client, key, 30)

		err := lock.AcquireWithTimeout(100*time.Millisecond, 10*time.Millisecond)
		assert.NoError(t, err)
	})

	t.Run("AcquireWithTimeout_Timeout", func(t *testing.T) {
		key := "test-lock-10"
		lock1 := NewRedisLockWithSeconds(client, key, 30)
		lock2 := NewRedisLockWithSeconds(client, key, 30)

		// Lock with lock1
		acquired, err := lock1.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired)

		// lock2 should time out
		err = lock2.AcquireWithTimeout(50*time.Millisecond, 10*time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, context.DeadlineExceeded, err)
	})

	t.Run("AcquireWithTimeout_WaitsForRelease", func(t *testing.T) {
		key := "test-lock-11"
		lock1 := NewRedisLockWithSeconds(client, key, 30)
		lock2 := NewRedisLockWithSeconds(client, key, 30)

		// Lock with lock1
		acquired, err := lock1.Acquire()
		assert.NoError(t, err)
		assert.True(t, acquired)

		// Start goroutine to release after delay
		go func() {
			time.Sleep(50 * time.Millisecond)
			released, err := lock1.Release()
			assert.NoError(t, err)
			assert.True(t, released)
		}()

		// lock2 should wait and eventually succeed
		err = lock2.AcquireWithTimeout(200*time.Millisecond, 10*time.Millisecond)
		assert.NoError(t, err)
	})
}
