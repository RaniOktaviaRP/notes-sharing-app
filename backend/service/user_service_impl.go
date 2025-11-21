package service

import (
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register (ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error) {
	tx, err := database.BeginTx(ctx)

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}
	
	user := domain.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: hashedPassword,
	}

	createdUser, err := s.userRepository.Create(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return web.UserResponse{}, err
	}

	tx.Commit()
	
	return web.UserResponse{
		Id:    createdUser.Id.String(),
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}, nil

	}