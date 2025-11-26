package helper

import (
	"database/sql"
	"notes-app/backend/model/domain"
	"notes-app/backend/model/web"
)

// Untuk throw panic jika ada error
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// Commit transaksi jika sukses, rollback jika gagal
func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}

// Convert domain.Note -> web.NoteResponse
func ToNoteResponse(note domain.Note) web.NoteResponse {
	return web.NoteResponse{
		Id:      note.Id,
		Title:   note.Title,
		Content: note.Content,
	}
}

func ToNoteResponses(notes []domain.Note) []web.NoteResponse {
	var responses []web.NoteResponse
	for _, note := range notes {
		responses = append(responses, ToNoteResponse(note))
	}
	return responses
}
