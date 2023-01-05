package bootstrap

import (
	"github.com/saxenashivang/api-gateway/api/controllers"
	"github.com/saxenashivang/api-gateway/api/middlewares"
	"github.com/saxenashivang/api-gateway/api/routes"
	"github.com/saxenashivang/api-gateway/client/grpc"
	"github.com/saxenashivang/api-gateway/infrastructure"
	"github.com/saxenashivang/api-gateway/lib"
	"go.uber.org/fx"
)

var Providers = fx.Options(
	grpc.Module,
	controllers.Module,
	routes.Module,
	infrastructure.Module,
	middlewares.Module,
	lib.Module,
)
