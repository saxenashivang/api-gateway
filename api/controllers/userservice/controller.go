package userservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"api-gateway/api/controllers/auth"
	"api-gateway/constants"
	userpb "api-gateway/services/user/proto"

	"api-gateway/lib"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
)

// UserController data type
type UserController struct {
	service        userpb.UserService
	logger         lib.Logger
	authController auth.JWTAuthService
}

// constructor to connect user microservice
func NewUserController(
	srv micro.Service,
	logger lib.Logger,
	authController auth.JWTAuthService,
) *UserController {
	c := userpb.NewUserService(constants.UserService, srv.Client())
	return &UserController{
		service:        c,
		logger:         logger,
		authController: authController,
	}
}

// GetUser gets one user
func (u *UserController) GetUser(c *gin.Context) {
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoidXNlciIsIklzc3VlciI6Im1vb2wiLCJVc2VybmFtZSI6InNoaXZhbmciLCJleHAiOjE2NzM3NzkwMDUsImlhdCI6MTY3MzUxOTgwNX0.grpsXyBjlDKcFAshJqfLyCnjK3k7xvpOWdBr776Va_U",
		"ID":            "1",
	})
	res, err := u.service.GetUser(ctx, &userpb.GetUserRequest{Id: "1"})
	fmt.Println(res, err)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res),
	})
}

// generate token will generate a signed token with claims
func (u *UserController) GenerateToken(c *gin.Context) {
	user := auth.User{
		ID:        12,
		Name:      "Shivang",
		Email:     "shivang@gmail.com",
		Age:       22,
		Role:      "Admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	token := u.authController.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "token generated successfully",
		"token":   token,
	})
}
