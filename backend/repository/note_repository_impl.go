package repository

import (
	"context"
	"database/sql"
	"time"
	"notes-app/backend/model/domain"
	"notes-app/backend/helper"
)

type NoteRepositoryImpl struct {}

func NewNoteRepositoryImpl(db *sql.DB) NoteRepository {
	return &NoteRepositoryImpl	{}
}

func (r *NoteRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note {
	SQL := "INSERT INTO notes (id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := tx.ExecContext(ctx, SQL, note.Id, note.Title, note.Content, time.Now(), time.Now())
	helper.PanicIfError(err)
	return note
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