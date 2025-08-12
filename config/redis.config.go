package config

import (
	"context"
	"net"

	"github.com/redis/go-redis/v9"
)

func NewRedisFailoverClient(cfg RedisConfig) *redis.Client {
	opt := &redis.FailoverOptions{
		MasterName:    cfg.MasterName,
		SentinelAddrs: cfg.SentinelAddrs,
		Password:      cfg.Password,
		DB:            cfg.DB,
		MaxRetries:    cfg.MaxRetries,
		DialTimeout:   cfg.DialTimeout,
		ReadTimeout:   cfg.ReadTimeout,
	}
	if cfg.RewriteAddr != nil {
		opt.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			addr = cfg.RewriteAddr(addr)
			var d net.Dialer
			d.Timeout = cfg.DialTimeout
			return d.DialContext(ctx, network, addr)
		}
	}
	return redis.NewFailoverClient(opt)
}
