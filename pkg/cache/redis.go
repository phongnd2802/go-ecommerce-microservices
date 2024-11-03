package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"
	"github.com/redis/go-redis/v9"
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

func NewRedisCache(cfg settings.RedisSetting) Cache {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})
	return &redisCache{
		r: rdb,
	}
}
