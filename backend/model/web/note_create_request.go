package web

import (
	"mime/multipart"
)

type NoteCreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Image *multipart.FileHeader `json:"-"`
}