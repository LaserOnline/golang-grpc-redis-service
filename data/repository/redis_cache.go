package repository

import (
	"context"
	"time"

	"github.com/auth-service/domain"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct{ c *redis.Client }

func NewRedisCache(c *redis.Client) domain.Cache {
	return &RedisCache{c: c}
}

func (r *RedisCache) Set(ctx context.Context, k string, v any) error {
	return r.c.Set(ctx, k, v, 0).Err()
}

func (r *RedisCache) Get(ctx context.Context, k string) (string, error) {
	return r.c.Get(ctx, k).Result()
}

func (r *RedisCache) Del(ctx context.Context, k string) (int64, error) {
	return r.c.Del(ctx, k).Result()
}

func (r *RedisCache) HSet(ctx context.Context, key string, fields map[string]interface{}) error {
	return r.c.HSet(ctx, key, fields).Err()
}

func (r *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.c.HGetAll(ctx, key).Result()
}

func (r *RedisCache) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return r.c.Expire(ctx, key, ttl).Result()
}

func (r *RedisCache) Close() error { return r.c.Close() }
