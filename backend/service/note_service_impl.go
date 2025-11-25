package service

import (
	"context"
	"database/sql"
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
	"notes-app/backend/repository"
	"os"
	"time"
)
type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	DB             *sql.DB
	Validate 	 *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, DB *sql.DB, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *NoteServiceImpl) Create(ctx context.Context, request web.NotesCreateRequest) (domain.Note, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note := domain.Note{
		Title:     request.Title,
		Content:   request.Content,
	}

	note = service.NoteRepository.Create(ctx, tx, note)
	return helper.ToNotesResponse(note)
}

func (service *NoteServiceImpl) Update(ctx context.Context, request web.NotesUpdateRequest) (domain.Note, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note := domain.Note{
		Id:        request.Id,
		Title:     request.Title,
		Content:   request.Content,
	}

	note = service.NoteRepository.Update(ctx, tx, note)
	return helper.ToNotesResponse(note)
}

func (service *NoteServiceImpl) Delete(ctx context.Context, request web.NotesDeleteRequest) error {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note := domain.Note{
		Id: request.Id,
	}

	service.NoteRepository.Delete(ctx, tx, note)
	return nil
}