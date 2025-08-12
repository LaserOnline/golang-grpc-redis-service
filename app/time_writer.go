package app

import (
	"context"

	"github.com/auth-service/domain"
)

type TimeWriter struct{ cache domain.Cache }

func NewTimeWriter(cache domain.Cache) *TimeWriter { return &TimeWriter{cache: cache} }

func (t *TimeWriter) Write(ctx context.Context, key, val string) error {
	return t.cache.Set(ctx, key, val)
}
func (t *TimeWriter) Read(ctx context.Context, key string) (string, error) {
	return t.cache.Get(ctx, key)
}
