package repository

import (
	"context"
	"database/sql"
	"time"
	"notes-app/backend/model/domain"
	"notes-app/backend/helper"
	"github.com/google/uuid"
)

type NoteRepositoryImpl struct {
	db *sql.DB
}

func NewNoteRepositoryImpl(db *sql.DB) NoteRepository {
	return &NoteRepositoryImpl{
		db: db,
	}
}

func (r *NoteRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, note domain.Note) (domain.Note, error) {

	if note.Id == uuid.Nil {
		note.Id = uuid.New()
	}

	now := time.Now()

	query := `
		INSERT INTO notes (id, title, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, content, created_at, updated_at
	`

	err := tx.QueryRowContext(
		ctx, query,
		note.Id, note.Title, note.Content, now, now).
		Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		return note, err
	}

	return note, nil
}

func (repository *NoteRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Note {
	rows, err := tx.QueryContext(ctx, "SELECT id, title, content FROM notes")
	helper.PanicIfError(err)
	defer rows.Close()

	var notes []domain.Note

	for rows.Next() {
		note := domain.Note{}
		err := rows.Scan(&note.Id, &note.Title, &note.Content)
		helper.PanicIfError(err)
		notes = append(notes, note)
	}

	return notes
}


func (r *NoteRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (domain.Note, error) {
	SQL := "SELECT id, title, content, created_at, updated_at FROM notes WHERE id = $1 AND deleted_at IS NULL"
	row := r.db.QueryRowContext(ctx, SQL, id)

	var note domain.Note
	err := row.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *NoteRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note {
	SQL := "UPDATE notes SET title = $1, content = $2, updated_at = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, SQL, note.Title, note.Content, time.Now(), note.Id)
	helper.PanicIfError(err)
	return note
}

func (r *NoteRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, note domain.Note) {
	SQL := "UPDATE notes SET deleted_at = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, time.Now(), note.Id)
	helper.PanicIfError(err)
}