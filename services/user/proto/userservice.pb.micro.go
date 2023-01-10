// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/userservice.proto

package userservice

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	// GetUserProfile return a profile of a user
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.CreateUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "UserService.DeleteUser", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	// GetUserProfile return a profile of a user
	GetUser(context.Context, *GetUserRequest, *User) error
	CreateUser(context.Context, *CreateUserRequest, *User) error
	DeleteUser(context.Context, *DeleteUserRequest, *emptypb.Empty) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		GetUser(ctx context.Context, in *GetUserRequest, out *User) error
		CreateUser(ctx context.Context, in *CreateUserRequest, out *User) error
		DeleteUser(ctx context.Context, in *DeleteUserRequest, out *emptypb.Empty) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *GetUserRequest, out *User) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}

func (h *userServiceHandler) CreateUser(ctx context.Context, in *CreateUserRequest, out *User) error {
	return h.UserServiceHandler.CreateUser(ctx, in, out)
}

func (h *userServiceHandler) DeleteUser(ctx context.Context, in *DeleteUserRequest, out *emptypb.Empty) error {
	return h.UserServiceHandler.DeleteUser(ctx, in, out)
}