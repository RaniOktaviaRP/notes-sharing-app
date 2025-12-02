package service

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"notes-app/backend/helper"
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
	"notes-app/backend/repository"
	"strings"

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

func processImage(file *multipart.FileHeader) ([]byte, error) {
	if file == nil {
		return nil, nil
	}

	filename := strings.ToLower(file.Filename)

	if !strings.HasSuffix(filename, ".jpg") &&
		!strings.HasSuffix(filename, ".jpeg") &&
		!strings.HasSuffix(filename, ".png") {
		return nil, fmt.Errorf("invalid image type: only JPG, JPEG, PNG allowed")
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	imageBytes, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}

func (service *NoteServiceImpl) Create(ctx context.Context, request web.NoteCreateRequest) (web.NoteResponse, error) {
	// validasi request
	if err := service.Validate.Struct(request); err != nil {
		return web.NoteResponse{}, err
	}

	// proses image tanpa panic
	imageBytes, err := processImage(request.Image)
	if err != nil {
		return web.NoteResponse{}, err
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note := domain.Note{
		Title:   request.Title,
		Content: request.Content,
		Image:   imageBytes,
	}

	note, err = service.NoteRepository.Create(ctx, tx, note)
	if err != nil {
		return web.NoteResponse{}, err
	}

	return helper.ToNoteResponse(note), nil
}

func (service *NoteServiceImpl) Update(ctx context.Context, request web.NoteUpdateRequest) (web.NoteResponse, error) {

	if err := service.Validate.Struct(request); err != nil {
		return web.NoteResponse{}, err
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note, err := service.NoteRepository.FindById(ctx, request.Id)
	if err != nil {
		return web.NoteResponse{}, err
	}

	if request.Title != nil {
		note.Title = *request.Title
	}

	if request.Content != nil {
		note.Content = *request.Content
	}

	// image optional
	if request.Image != nil {
		imageBytes, err := processImage(request.Image)
		if err != nil {
			return web.NoteResponse{}, err
		}
		note.Image = imageBytes
	}

	note, err = service.NoteRepository.Update(ctx, tx, note)
	if err != nil {
		return web.NoteResponse{}, err
	}

	return helper.ToNoteResponse(note), nil
}

func (service *NoteServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note, err := service.NoteRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	return service.NoteRepository.Delete(ctx, tx, note)
}

func (service *NoteServiceImpl) GetAll(ctx context.Context) ([]web.NoteResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes, err := service.NoteRepository.GetAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return helper.ToNoteResponses(notes), nil
}

func (service *NoteServiceImpl) FindById(ctx context.Context, id uuid.UUID) (web.NoteResponse, error) {
	note, err := service.NoteRepository.FindById(ctx, id)
	if err != nil {
		return web.NoteResponse{}, err
	}

	return helper.ToNoteResponse(note), nil
}
