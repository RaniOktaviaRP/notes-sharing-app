package service

import (
	"context"
	"database/sql"
	"notes-app/backend/helper"
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
	"notes-app/backend/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, db *sql.DB, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *NoteServiceImpl) Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note := domain.Note{
		Title:   request.Title,
		Content: request.Content,
	}

	note, err = service.NoteRepository.Create(ctx, tx, note) 
	helper.PanicIfError(err)

	return helper.ToNoteResponse(note)
}

func (service *NoteServiceImpl) Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note, err := service.NoteRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	note.Title = *request.Title
	note.Content = *request.Content

	note = service.NoteRepository.Update(ctx, tx, note)
	return helper.ToNoteResponse(note)
}

func (service *NoteServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note, err := service.NoteRepository.FindById(ctx, id)
	helper.PanicIfError(err)

	service.NoteRepository.Delete(ctx, tx, note)
	 return nil
}

func (service *NoteServiceImpl) GetAll(ctx context.Context) []web.NoteResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes := service.NoteRepository.GetAll(ctx, tx)
	return helper.ToNoteResponses(notes)
}

func (service *NoteServiceImpl) FindById(ctx context.Context, id uuid.UUID) web.NoteResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    note, err := service.NoteRepository.FindById(ctx, id)
    helper.PanicIfError(err)

    return helper.ToNoteResponse(note)
}

