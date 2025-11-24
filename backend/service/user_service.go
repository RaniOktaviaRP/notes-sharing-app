package service

import (
	"context"
	"notes-app/backend/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error)
	Login(ctx context.Context, request web.UserLoginRequest) (string, error)
}