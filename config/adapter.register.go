package config

import (
	"context"
	"fmt"

	"github.com/auth-service/app"
	"github.com/auth-service/data/repository"
	"github.com/auth-service/domain"
	"github.com/redis/go-redis/v9"
)

type AppContainer struct {
	RedisClient *redis.Client
	Cache       domain.Cache
	Service     app.IRedisService
	Close       func() error
}

func BuildApp() (*AppContainer, error) {
	rcfg := LoadRedisConfig()

	rdb := NewRedisFailoverClient(rcfg)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	cache := repository.NewRedisCache(rdb)
	svc := app.NewRedisService(cache)

	return &AppContainer{
		RedisClient: rdb,
		Cache:       cache,
		Service:     svc,
		Close:       rdb.Close,
	}, nil
}
