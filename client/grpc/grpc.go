package grpc

import (
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewGrpcConnection),
)

func NewGrpcConnection() micro.Service {
	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	srv.Init()
	return srv
}
