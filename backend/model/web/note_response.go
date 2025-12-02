package web 

import (
		"github.com/google/uuid"
)

type NoteResponse struct {
	Id 		uuid.UUID   "json:id"
	Title 	string 		"json:title"
	Content string		"json:content"
	Image string`json:"-"`
}