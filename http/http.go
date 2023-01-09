package http

import (
	"github.com/saxenashivang/api-gateway/http/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(router.NewRouter),
)
