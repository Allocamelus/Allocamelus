package stores

import (
	"context"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/pkg/fiberutil/session"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/go-redis/redis/v8"
	"k8s.io/klog/v2"
)

// ErrNilKey returned when key has 0 or less length
var ErrNilKey = errors.New("stores.redis: Error 0 length key")

// Redis struct
type Redis struct {
	db *redis.Client
}

// NewRedis Storage
func NewRedis(redis *redis.Client) *Redis {
	// Check Redis
	if err := redis.Ping(context.Background()).Err(); err != nil {
		klog.Fatal(err)
	}
	return &Redis{db: redis}
}

// Get []bytes
func (r *Redis) Get(key string) (*session.Session, error) {
	if len(key) > 0 {
		data, err := r.db.Get(context.Background(), key).Bytes()
		if err != nil {
			if err != redis.Nil {
				logger.Error(err)
			}
			return &session.Session{}, session.ErrNoSession
		}
		s := new(session.Session)
		_, err = s.UnmarshalMsg(data)
		if err != nil {
			logger.Error(err)
			return &session.Session{}, session.ErrNoSession
		}
		return s, nil
	}
	return &session.Session{}, ErrNilKey
}

// Set []bytes
func (r *Redis) Set(s *session.Session) error {
	if len(s.Key) > 0 {
		data, err := s.MarshalMsg(nil)
		// This Error should not happen
		logger.Error(err)
		return r.db.Set(context.Background(), s.Key, data, time.Until(s.Expires)).Err()
	}
	return ErrNilKey
}

// Delete key
func (r *Redis) Delete(key string) error {
	if len(key) > 0 {
		return r.db.Del(context.Background(), key).Err()
	}
	return ErrNilKey
}
