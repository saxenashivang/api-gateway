package routes

import (
	rest "github.com/saxenashivang/api-gateway/infrastructure"
	"github.com/saxenashivang/api-gateway/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger  lib.Logger
	handler rest.Router
}

func NewUserRoutes(
	logger lib.Logger,
	handler rest.Router,
) *UserRoutes {
	return &UserRoutes{
		logger:  logger,
		handler: handler,
	}
}

// Setup user routes
func (s *UserRoutes) Setup() {
	s.logger.Info("Setting up routes")

	api := s.handler.Group("/api").Use()

	// api.GET("/user", s.PaginationMiddleware.Handle(), s.userController.GetUser)
	// api.GET("/user/:id", s.userController.GetOneUser)
	// api.POST("/user", s.userController.SaveUser)
	// api.PUT("/user/:id",
	// 	s.uploadMiddleware.Push(s.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
	// 	s.userController.UpdateUser,
	// )
	api.DELETE("/user/:id")

}
