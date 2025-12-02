package web

import (
		"mime/multipart"
		"github.com/google/uuid"
)

type NoteUpdateRequest struct {
	Id      uuid.UUID `json:"id"` 
	Title   *string   `json:"title"`   
	Content *string   `json:"content"` 
	Image *multipart.FileHeader `json:"-"`
}
