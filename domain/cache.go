package domain

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) (int64, error)

	HSet(ctx context.Context, key string, fields map[string]interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Expire(ctx context.Context, key string, ttl time.Duration) (bool, error)

	Close() error
}
