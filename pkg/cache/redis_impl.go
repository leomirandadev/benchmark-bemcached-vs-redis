package cache

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

type redisImpl struct {
	cache *redis.Client
}

func NewRedis(host string) ICache {
	return &redisImpl{
		cache: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: "",
			DB:       0,
		}),
	}
}

func (c *redisImpl) Get(ctx context.Context, key string) (interface{}, error) {
	cmd := c.cache.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	return cmd.Val(), nil
}

func (c *redisImpl) Set(ctx context.Context, key string, value []byte) error {
	cmd := c.cache.Set(ctx, key, value, expirationTime)
	return cmd.Err()
}
