package routes

import (
	"github.com/saxenashivang/api-gateway/api/controllers"
	"github.com/saxenashivang/api-gateway/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController *controllers.UserController
}

func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController *controllers.UserController,
) *UserRoutes {
	return &UserRoutes{
		userController: userController,
		logger:         logger,
		handler:        handler,
	}
}

// Setup user routes
func (s *UserRoutes) Setup() {
	s.logger.Info("Setting up user routes")

	api := s.handler.Gin.Group("/api").Use()
	api.GET("/user/:id", s.userController.GetUser)

}
