package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"notes-app/backend/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error
}