package handler

import (
	"context"
	"time"

	"github.com/auth-service/app"
	cachev1 "github.com/auth-service/stud"
	"github.com/auth-service/validation"
)

type CacheHandler struct {
	svc       app.IRedisService
	validator validation.SetValidator
	cachev1.UnimplementedCacheServiceServer
}

func NewCacheHandler(svc app.IRedisService, v validation.SetValidator) *CacheHandler {
	return &CacheHandler{svc: svc, validator: v}
}

// Function Set Data Redis
func (h *CacheHandler) SetData(ctx context.Context, in *cachev1.SetRequest) (*cachev1.SetReply, error) {
	if err := h.validator.ValidateSet(in.GetKey(), in.GetValue()); err != nil {
		return &cachev1.SetReply{Status: false, Message: err.Error()}, nil
	}
	if err := h.svc.RedisSetData(ctx, in.GetKey(), in.GetValue()); err != nil {
		return &cachev1.SetReply{Status: false, Message: err.Error()}, nil
	}
	return &cachev1.SetReply{Status: true, Message: "successfully"}, nil
}

// Function Get Data Redis By Key
func (h *CacheHandler) GetData(ctx context.Context, in *cachev1.GetRequest) (*cachev1.GetReply, error) {
	if err := h.validator.ValidateGet(in.GetKey()); err != nil {
		return &cachev1.GetReply{Status: false, Message: err.Error()}, nil
	}
	val, err := h.svc.RedisGetData(ctx, in.GetKey())
	if err != nil {
		return &cachev1.GetReply{Status: false, Key: in.GetKey(), Message: err.Error()}, nil
	}
	return &cachev1.GetReply{Status: true, Key: in.GetKey(), Value: val, Message: "successfully"}, nil
}

// Function Delete Data Redis By Key
func (h *CacheHandler) DelData(ctx context.Context, in *cachev1.DelRequest) (*cachev1.DelReply, error) {
	if err := h.validator.ValidateGet(in.GetKey()); err != nil {
		return &cachev1.DelReply{Status: false, Message: err.Error()}, nil
	}
	if err := h.svc.RedisDelData(ctx, in.GetKey()); err != nil {
		return &cachev1.DelReply{Status: false, Message: err.Error()}, nil
	}
	return &cachev1.DelReply{Status: true, Message: "successfully deleted"}, nil
}

// Function Add Session By Key

func (h *CacheHandler) HsetSession(ctx context.Context, in *cachev1.SetSession) (*cachev1.SetSessionReply, error) {
	if err := h.validator.ValidateGet(in.GetKey()); err != nil {
		return &cachev1.SetSessionReply{Status: false, Message: err.Error()}, nil
	}
	if err := h.validator.ValidateGet(in.GetUuid()); err != nil {
		return &cachev1.SetSessionReply{Status: false, Message: err.Error()}, nil
	}
	if err := h.validator.ValidateGet(in.GetSession()); err != nil {
		return &cachev1.SetSessionReply{Status: false, Message: err.Error()}, nil
	}
	if in.GetTtlSeconds() <= 0 {
		return &cachev1.SetSessionReply{Status: false, Message: "invalid ttl_seconds"}, nil
	}

	fields := map[string]any{
		"uuid":    in.GetUuid(),
		"session": in.GetSession(),
	}
	if err := h.svc.RedisHsetDataWithTTL(ctx, in.GetKey(), fields, time.Duration(in.GetTtlSeconds())*time.Second); err != nil {
		return &cachev1.SetSessionReply{Status: false, Message: err.Error()}, nil
	}
	return &cachev1.SetSessionReply{Status: true, Message: "session stored successfully"}, nil
}

// Function Get Session By Key

func (h *CacheHandler) HgetSession(ctx context.Context, in *cachev1.GetSessionRequest) (*cachev1.GetSessionReply, error) {
	if err := h.validator.ValidateGet(in.GetKey()); err != nil {
		return &cachev1.GetSessionReply{Status: false, Message: err.Error()}, nil
	}

	m, err := h.svc.RedisHgetAll(ctx, in.GetKey())
	if err != nil {
		return &cachev1.GetSessionReply{Status: false, Message: err.Error()}, nil
	}
	if len(m) == 0 {
		return &cachev1.GetSessionReply{
			Status:  false,
			Message: "hash not found or empty",
			Key:     in.GetKey(),
		}, nil
	}

	return &cachev1.GetSessionReply{
		Status:  true,
		Message: "ok",
		Key:     in.GetKey(),
		Uuid:    m["uuid"],
		Session: m["session"],
	}, nil
}
