package repository

import (
	"context"
	"database/sql"
	"errors"
	"notes-app/backend/model/domain"
	"github.com/google/uuid"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := `INSERT INTO users (email, name, password_hash, created_at, updated_at) 
	        VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, email, name, password_hash, created_at, updated_at, deleted_at`
	err := tx.QueryRowContext(ctx, SQL, user.Email, user.Name, user.PasswordHash).Scan(
		&user.Id, &user.Email, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, email, name, password_hash, created_at, updated_at, deleted_at FROM users WHERE email = $1"
	row := tx.QueryRowContext(ctx, SQL, email)
	user := domain.User{}
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (domain.User, error) {
	SQL := "SELECT id, email, name, password_hash, created_at, updated_at, deleted_at FROM users WHERE id = $1"
	row := tx.QueryRowContext(ctx, SQL, id)
	user := domain.User{}
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := `UPDATE users 
	        SET email = $1, name = $2, password_hash = $3, updated_at = NOW() 
	        WHERE id = $4 
	        RETURNING id, email, name, password_hash, created_at, updated_at, deleted_at`
	row := tx.QueryRowContext(ctx, SQL, user.Email, user.Name, user.PasswordHash, user.Id)
	updatedUser := domain.User{}
	err := row.Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Name, &updatedUser.PasswordHash,
		&updatedUser.CreatedAt, &updatedUser.UpdatedAt, &updatedUser.DeletedAt)
	if err != nil {
		return domain.User{}, err
	}
	return updatedUser, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	SQL := "UPDATE users SET deleted_at = NOW() WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, id)
	return err
}