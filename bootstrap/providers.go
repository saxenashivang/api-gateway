package bootstrap

import (
	"github.com/saxenashivang/api-gateway/api/middlewares"
	"github.com/saxenashivang/api-gateway/api/routes"
	"github.com/saxenashivang/api-gateway/infrastructure"
	"github.com/saxenashivang/api-gateway/lib"
	"go.uber.org/fx"
)

func GetProviderOptions() []fx.Option {
	return []fx.Option{
		// controllers.Module,
		routes.Module,
		infrastructure.Module,
		middlewares.Module,
		lib.Module,
	}
}
