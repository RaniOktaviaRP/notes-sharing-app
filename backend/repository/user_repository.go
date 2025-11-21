package repository	

import (
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
)	

type UserRepository interface {
	Create(request web.UserRegisterRequest) (web.UserResponse, error)
	FindByEmail(email string) (web.UserResponse, error)
	FindById(id string) (web.UserResponse, error)
	Update(user web.UserResponse) (web.UserResponse, error)
	Delete(id string) error
}