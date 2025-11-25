package web

type NotesCreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}