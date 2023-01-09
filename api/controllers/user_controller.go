package controllers

import (
	"context"
	"fmt"
	"net/http"
	userpb "userservice/proto"

	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/api-gateway/constants"
	"github.com/saxenashivang/api-gateway/lib"
	"go-micro.dev/v4"
)

// UserController data type
type UserController struct {
	service userpb.UserService
	logger  lib.Logger
}

// TODO - figure out name resulution conyainer
func NewUserController(srv micro.Service, logger lib.Logger) *UserController {
	c := userpb.NewUserService(constants.UserService, srv.Client())
	return &UserController{
		service: c,
		logger:  logger,
	}
}

// GetOneUser gets one user
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
