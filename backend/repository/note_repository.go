package repository

import (
	"context"
	"database/sql"
	"notes-app/backend/model/domain"

	"github.com/google/uuid"
)

type NoteRepository interface {
	Create(ctx context.Context, tx *sql.Tx, note domain.Note) (domain.Note, error)
	Update(ctx context.Context, tx *sql.Tx, note domain.Note) (domain.Note, error)
	FindById(ctx context.Context, id uuid.UUID) (domain.Note, error)
	Delete(ctx context.Context, tx *sql.Tx, note domain.Note) error
	GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Note, error)
}
