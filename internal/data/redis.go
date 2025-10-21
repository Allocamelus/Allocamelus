package data

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (d *Data) initRedis() {
	d.redis = redis.NewClient(&redis.Options{
		Addr:     d.Config.Redis.Host,
		Username: d.Config.Redis.User,
		Password: d.Config.Redis.Password, // no password set
	})
}

// Set stores values in redis store
func (d *Data) Set(key string, value []byte, expiration time.Duration) error {
	if key == "" || len(value) == 0 {
		return nil
	}

	return d.redis.Set(context.Background(), key, value, expiration).Err()
}

// Get retrieves values from redis store
// Default returns nil, nil
func (d *Data) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, nil
	}

	value, err := d.redis.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}

	return value, err
}

// Delete deletes a key from redis store
func (d *Data) Delete(key string) error {
	if key == "" {
		return nil
	}

	return d.redis.Del(context.Background(), key).Err()
}

// Delete deletes keys from redis store
func (d *Data) DeleteKeys(keys ...string) error {
	return d.redis.Del(context.Background(), keys...).Err()
}

// Reset redis store
func (d *Data) Reset() error {
	return d.redis.FlushDB(context.Background()).Err()
}

// Close redis store
func (d *Data) Close() error {
	d.database.Close()
	return d.redis.Close()
}
