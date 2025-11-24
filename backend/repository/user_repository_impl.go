package repository 

import (
	"context"
	"database/sql"
	"errors"
	"notes-app/backend/helper"
	"github.com/google/uuid"
	"notes-app/backend/model/domain"
)

type UserRepositoryImpl struct {}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (email, name, password_hash) VALUES ($1, $2, $3) RETURNING id, email, name"
	err := tx.QueryRowContext(ctx, SQL, user.Email, user.Name, user.PasswordHash).Scan(&user.Id)
	helper.PanicIfError(err)
	return user
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, email, name, password_hash FROM users WHERE email = $1"
	row := tx.QueryRowContext(ctx, SQL, email)
	FindByEmailUser := domain.User{}
	err := row.Scan(&FindByEmailUser.Id, &FindByEmailUser.Email, &FindByEmailUser.Name, &FindByEmailUser.PasswordHash)
	if err != nil {
		return domain.User{}, err
	}
	return FindByEmailUser, nil
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (domain.User, error) {
	SQL := "SELECT id, email, name, password_hash FROM users WHERE id = $1"
	row := tx.QueryRowContext(ctx, SQL, id)
	FindUser := domain.User{}
	err := row.Scan(&FindUser.Id, &FindUser.Email, &FindUser.Name, &FindUser.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return FindUser, errors.New("user not found")
		}
		return FindUser, err
	}
	return FindUser, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "UPDATE users SET email = $1, name = $2, password_hash = $3 WHERE id = $4 RETURNING id, email, name"
	row := tx.QueryRowContext(ctx, SQL, user.Email, user.Name, user.PasswordHash, user.Id)
	UpdateUser := domain.User{}
	err := row.Scan(&UpdateUser.Id, &UpdateUser.Email, &UpdateUser.Name)
	if err != nil {
		return domain.User{}, err
	}
	return UpdateUser, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	SQL := "UPDATE users SET deleted_at = NOW() WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}
