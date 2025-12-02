package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"notes-app/backend/model/domain"

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
		INSERT INTO notes (id, title, content, image, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, content, image, created_at, updated_at
	`
	err := tx.QueryRowContext(ctx, query,
		note.Id, note.Title, note.Content, note.Image, now, now).
		Scan(&note.Id, &note.Title, &note.Content,&note.Image, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *NoteRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Note, error) {
	rows, err := tx.QueryContext(ctx, "SELECT id, title, content, image, created_at, updated_at FROM notes WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []domain.Note

	for rows.Next() {
		note := domain.Note{}
		err := rows.Scan(&note.Id, &note.Title, &note.Content,&note.Image, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *NoteRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (domain.Note, error) {
	const SQL = `
		SELECT id, title, content, image, created_at, updated_at
		FROM notes
		WHERE id = $1 AND deleted_at IS NULL
	`

	var note domain.Note
	row := r.db.QueryRowContext(ctx, SQL, id)
	err := row.Scan(&note.Id, &note.Title, &note.Content,&note.Image, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return note, fmt.Errorf("note with id %s not found", id.String())
		}
		return note, err
	}

	return note, nil
}

func (r *NoteRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, note domain.Note) (domain.Note, error) {
	SQL := "UPDATE notes SET title = $1, content = $2, image = $3, updated_at = $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, note.Title, note.Content, note.Image, time.Now(), note.Id)

	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *NoteRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, note domain.Note) error {
	SQL := "UPDATE notes SET deleted_at = NOW() WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, note.Id)
	return err
}
