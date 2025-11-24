package service

import (
	"os"
	"context"
	"database/sql"
	"notes-app/backend/helper"
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
	"github.com/golang-jwt/jwt/v5"
	"notes-app/backend/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB 		   *sql.DB	
}

func NewUserServiceImpl(UserRepository repository.UserRepository, db *sql.DB) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		DB:			 db,
	}
}

func (s *UserServiceImpl) Register (ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer tx.Rollback()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	
	user := domain.User{
		Email:    request.Email,
		Name:     request.Name,
		PasswordHash: string(hashedPassword),
	}

	createdUser := s.UserRepository.Create(ctx, tx, user)
	
	return web.UserResponse{
		Id:    createdUser.Id.String(),
		Email: createdUser.Email,
		Name:  createdUser.Name,
	},nil

	}

// generate jwt token
func (s *UserServiceImpl) Login (ctx context.Context, request web.UserLoginRequest) (string, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return "", err
	}

	user, err := s.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		tx.Rollback()
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return tokenString, nil
}