package service

import (
	"notes-app/backend/model/web"
)

type UserService interface {
	Register(request web.UserRegisterRequest) (web.UserResponse, error)
}