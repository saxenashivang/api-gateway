package userservice

import (
	"context"
	"fmt"
	"net/http"

	"api-gateway/constants"
	userpb "api-gateway/services/user/proto"

	"api-gateway/lib"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

// UserController data type
type UserController struct {
	service userpb.UserService
	logger  lib.Logger
}

// constructor to connect user microservice
func NewUserController(srv micro.Service, logger lib.Logger) *UserController {
	c := userpb.NewUserService(constants.UserService, srv.Client())
	return &UserController{
		service: c,
		logger:  logger,
	}
}

// GetUser gets one user
func (u *UserController) GetUser(c *gin.Context) {
	res, err := u.service.GetUser(context.Background(), &userpb.GetUserRequest{Id: "1"})
	fmt.Println(res, err)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res),
	})
}
