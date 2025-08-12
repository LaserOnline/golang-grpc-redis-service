package main

import (
	"log"

	"github.com/auth-service/config"
	"github.com/auth-service/presentation/grpc/handler"
	cachev1 "github.com/auth-service/stud"
	"github.com/auth-service/validation"
	"google.golang.org/grpc"
)

func main() {

	c, err := config.BuildApp()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	v := validation.NewSimpleSetValidator()
	cacheHandler := handler.NewCacheHandler(c.Service, v)

	gcfg := config.LoadGRPCConfig()
	s, lis, err := config.NewGRPCServer(
		gcfg,
		func(gs *grpc.Server) {
			cachev1.RegisterCacheServiceServer(gs, cacheHandler)
		},
	)
	if err != nil {
		log.Fatalf("grpc build failed: %v", err)
	}

	log.Printf("gRPC listening on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve failed: %v", err)
	}
}
