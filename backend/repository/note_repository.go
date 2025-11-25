package repository

import (
	"context"
	"database/sql"
	"notes-app/backend/model/domain"
)

type NoteRepository interface {
	Create(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note
	Update(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note
	Delete(ctx context.Context, tx *sql.Tx, note domain.Note)
}
