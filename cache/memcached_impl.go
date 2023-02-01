package cache

import (
	"context"

	"github.com/bradfitz/gomemcache/memcache"
)

type memcachedImpl struct {
	cache *memcache.Client
}

func NewMemcached(host string) ICache {
	client := memcache.New(host)

	return &memcachedImpl{
		cache: client,
	}
}

func (c *memcachedImpl) Get(ctx context.Context, key string) (interface{}, error) {
	item, err := c.cache.Get(key)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *memcachedImpl) Set(ctx context.Context, key string, value []byte) error {
	err := c.cache.Set(&memcache.Item{Key: key, Value: value, Expiration: int32(expirationTime.Seconds())})
	if err != nil {
		return err
	}
	return nil
}
