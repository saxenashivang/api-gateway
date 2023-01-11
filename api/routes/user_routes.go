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
}

func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController *userservice.UserController,
	rateLimit middlewares.RateLimitMiddleware,
) *UserRoutes {
	return &UserRoutes{
		userController:      userController,
		logger:              logger,
		handler:             handler,
		rateLimitMiddleware: rateLimit,
	}
}

// Setup user routes
func (s *UserRoutes) Setup() {
	s.logger.Info("Setting up user routes")

	api := s.handler.Gin.Group("/api").Use(
		s.rateLimitMiddleware.Handle())
	api.GET("/user/:id", s.userController.GetUser)

}
