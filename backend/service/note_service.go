package service

import (
	"context"
	"notes-app/backend/model/web"
	"github.com/google/uuid"
)

type NoteService interface {
	Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse
	Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) []web.NoteResponse
	FindById(ctx context.Context, id uuid.UUID) web.NoteResponse
}
