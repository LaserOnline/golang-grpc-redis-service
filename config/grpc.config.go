package config

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCConfig struct {
	Host string
	Port string
	Opts []grpc.ServerOption
}

type Registrar func(s *grpc.Server)

func LoadGRPCConfig() GRPCConfig {
	LoadDotenv()
	p := ActiveProfile()
	host := getp(p, "GOLANG_SERVICE_ADDRESS", "127.0.0.1")
	port := getp(p, "GOLANG_SERVICE_PORT", "50051")
	return GRPCConfig{Host: host, Port: port}
}

func NewGRPCServer(cfg GRPCConfig, register Registrar) (*grpc.Server, net.Listener, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	s := grpc.NewServer(cfg.Opts...)
	reflection.Register(s)

	if register != nil {
		register(s)
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil, err
	}
	return s, lis, nil
}
