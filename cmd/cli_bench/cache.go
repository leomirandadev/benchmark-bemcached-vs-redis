package main

import (
	"bench-memcached-vs-redis/pkg/cache"
	"context"
	"log"
	"time"
)

func cacheCall(host, provider, action string) time.Duration {

	var cacheSystem cache.ICache

	switch provider {
	case "memcached":
		cacheSystem = cache.NewMemcached(host)
	case "redis":
		cacheSystem = cache.NewRedis(host)
	default:
		log.Fatal("error to identify provider", provider)
	}

	data := []byte("value")

	now := time.Now()
	switch action {
	case "get":
		_, err := cacheSystem.Get(context.Background(), "key")
		if err != nil {
			log.Fatal(err)
		}
		return time.Since(now)
	case "set":
		err := cacheSystem.Set(context.Background(), "key", data)
		if err != nil {
			log.Fatal(err)
		}
		return time.Since(now)
	default:
		log.Fatal("error to identify action", action)
	}

	return 0
}
