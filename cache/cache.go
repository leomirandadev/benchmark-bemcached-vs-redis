package cache

import (
	"context"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value []byte) error
}

const expirationTime = time.Duration(time.Minute)
