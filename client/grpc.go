package client

import (
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go.uber.org/fx"
)

// Module exported for initializing application
var GRPCModule = fx.Options(
	fx.Provide(newGRPCClientConnection),
)

// private function - create a grpc client using
func newGRPCClientConnection() micro.Service {
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	service.Init()
	return service
}
