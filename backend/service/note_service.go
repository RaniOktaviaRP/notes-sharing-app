package service

import (
	"context"
	"notes-app/backend/model/web"

	"github.com/google/uuid"
)

type NoteService interface {
	Create(ctx context.Context, request web.NoteCreateRequest) (web.NoteResponse, error)
	Update(ctx context.Context, request web.NoteUpdateRequest) (web.NoteResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]web.NoteResponse, error)
	FindById(ctx context.Context, id uuid.UUID) (web.NoteResponse, error)
}
