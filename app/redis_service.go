package app

import (
	"context"
	"fmt"
	"time"

	"github.com/auth-service/domain"
)

type IRedisService interface {
	RedisSetData(ctx context.Context, key, value string) error
	RedisGetData(ctx context.Context, key string) (string, error)
	RedisDelData(ctx context.Context, key string) error

	RedisHsetData(ctx context.Context, key string, fields map[string]interface{}) error
	RedisHgetAll(ctx context.Context, key string) (map[string]string, error)
	RedisHsetDataWithTTL(ctx context.Context, key string, fields map[string]interface{}, ttl time.Duration) error
}

type RedisService struct {
	cache domain.Cache
}

func NewRedisService(cache domain.Cache) *RedisService { return &RedisService{cache: cache} }

func (s *RedisService) RedisSetData(ctx context.Context, key, value string) error {
	return s.cache.Set(ctx, key, value)
}

func (s *RedisService) RedisGetData(ctx context.Context, key string) (string, error) {
	return s.cache.Get(ctx, key)
}

func (s *RedisService) RedisDelData(ctx context.Context, key string) error {
	n, err := s.cache.Del(ctx, key)
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("key not found")
	}
	return nil
}

func (s *RedisService) RedisHsetData(ctx context.Context, key string, fields map[string]interface{}) error {
	return s.cache.HSet(ctx, key, fields)
}

func (s *RedisService) RedisHgetAll(ctx context.Context, key string) (map[string]string, error) {
	return s.cache.HGetAll(ctx, key)
}

func (s *RedisService) RedisHsetDataWithTTL(ctx context.Context, key string, fields map[string]interface{}, ttl time.Duration) error {
	if err := s.cache.HSet(ctx, key, fields); err != nil {
		return err
	}
	if ttl > 0 {
		ok, err := s.cache.Expire(ctx, key, ttl)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("failed to set TTL")
		}
	}
	return nil
}
