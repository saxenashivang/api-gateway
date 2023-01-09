package routes

import (
	"github.com/saxenashivang/api-gateway/api/controllers"
	http "github.com/saxenashivang/api-gateway/http/router"
	"github.com/saxenashivang/api-gateway/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        http.Router
	userController *controllers.UserController
}

func NewUserRoutes(
	logger lib.Logger,
	handler http.Router,
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
	s.logger.Info("Setting up routes")

	api := s.handler.Group("/api").Use()
	api.GET("/user/:id", s.userController.GetUser)

}
