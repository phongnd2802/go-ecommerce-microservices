package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisCache struct {
	r *redis.Client
}

// Get implements Cache.
func (redis *redisCache) Get(ctx context.Context, key string) (string, error) {
	return redis.r.Get(ctx, key).Result()
}

// Set implements Cache.
func (redis *redisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return redis.r.Set(ctx, key, value, expiration).Err()
}

func NewRedisCache(addr string) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})
	return &redisCache{
		r: rdb,
	}
}
