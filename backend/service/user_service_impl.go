package service

import (
	"context"
	"database/sql"
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
	"notes-app/backend/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserServiceImpl(UserRepository repository.UserRepository, db *sql.DB) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		DB:             db,
	}
}

func (s *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer tx.Rollback()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := domain.User{
		Email:        request.Email,
		Name:         request.Name,
		PasswordHash: string(hashedPassword),
	}

	createdUser, err := s.UserRepository.Create(ctx, tx, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:    createdUser.Id.String(),
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (string, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	user, err := s.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": user.Id.String(),
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
