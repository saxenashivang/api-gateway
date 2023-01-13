package routes

import (
	"api-gateway/api/controllers/userservice"
	"api-gateway/api/middlewares"
	"api-gateway/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger              lib.Logger
	handler             lib.RequestHandler
	userController      *userservice.UserController
	rateLimitMiddleware middlewares.RateLimitMiddleware
	authMiddleware      middlewares.JWTAuthMiddleware
}

func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController *userservice.UserController,
	rateLimit middlewares.RateLimitMiddleware,
	authMiddleware middlewares.JWTAuthMiddleware,
) *UserRoutes {
	return &UserRoutes{
		userController:      userController,
		logger:              logger,
		handler:             handler,
		rateLimitMiddleware: rateLimit,
		authMiddleware:      authMiddleware,
	}
}

// Setup user routes
func (s *UserRoutes) Setup() {
	s.logger.Info("Setting up user routes")

	api := s.handler.Gin.Group("/api").Use(
		s.rateLimitMiddleware.Handle(), s.authMiddleware.Handler())
	api.GET("/user/:id", s.userController.GetUser)

	// TODO : temporary route
	s.handler.Gin.GET("/generate-token", s.userController.GenerateToken)

}
