package repository	

import (
	"github.com/google/uuid"
	"notes-app/backend/model/domain"
	"context"
	"database/sql"
)	

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, User domain.User) domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, User domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error
}