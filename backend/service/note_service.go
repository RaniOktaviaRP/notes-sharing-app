package seervice

import (
	"context"
	"notes-app/backend/model/web"
)

type NoteService interface {
	Create(ctx context.Context, request web.NotesCreateRequest) web.NotesResponse
	Update(ctx context.Context, request web.NotesUpdateRequest) (domain.Note, error)
	Delete(ctx context.Context, request web.NotesDeleteRequest) error
}
