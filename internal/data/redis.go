package data

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"k8s.io/klog/v2"
)

func (d *Data) initRedis() {
	d.redis = redis.NewClient(&redis.Options{
		Addr:     d.Config.Redis.Host,
		Username: d.Config.Redis.User,
		Password: d.Config.Redis.Password, // no password set
	})
}

// Set stores values in redis store
func (d *Data) Set(key string, value interface{}, expiration time.Duration) {
	ctx := context.Background()
	if err := d.redis.Set(ctx, key, value, expiration).Err(); err != nil {
		klog.Error(err)
	}
}

// Get retrieves values from redis store
// Default returns nil
// Interface panics if type hinted for anything but string
func (d *Data) Get(key string) *redis.StringCmd {
	if len(key) <= 0 {
		return &redis.StringCmd{}
	}
	ctx := context.Background()
	return d.redis.Get(ctx, key)
}

// Delete deletes values from redis store
func (d *Data) Delete(keys ...string) {
	ctx := context.Background()
	d.redis.Del(ctx, keys...)
}
