package controllers

import (
	"api-gateway/api/controllers/auth"
	"api-gateway/api/controllers/userservice"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(userservice.NewUserController),
	fx.Provide(auth.NewJWTAuthService),
)
