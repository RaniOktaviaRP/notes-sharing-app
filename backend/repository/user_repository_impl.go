package repository 

import (
	"notes-app/backend/model"
	"notes-app/backend/domain"
)

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3, $4) RETURNING id, email, name"
	row := tx.QueryRowContext(ctx, query, user.Id, user.Email, user.Name, user.Password)
	var createdUser domain.User
	err := row.Scan(&createdUser.Id, &createdUser.Email, &createdUser.Name)
	if err != nil {
		return domain.User{}, err
	}
	return createdUser, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	query := "SELECT id, email, name, password FROM users WHERE email = $1"
	row := tx.QueryRowContext(ctx, query, email)
	var user domain.User
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindById(ctx context.Context, tx *sql.Tx, id string) (domain.User, error) {
	query := "SELECT id, email, name, password FROM users WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)
	var user domain.User
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "UPDATE users SET email = $1, name = $2, password = $3 WHERE id = $4 RETURNING id, email, name"
	row := tx.QueryRowContext(ctx, query, user.Email, user.Name, user.Password, user.Id)
	var user domain.User
	err := row.Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	return err
}
