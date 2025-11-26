package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/google/uuid"
	"notes-app/backend/helper"
	"notes-app/backend/model/web"
	"notes-app/backend/service"
)

type NoteController struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return NoteController{
		NoteService: noteService,
	}
}

// Create godoc
// @Summary Create a new note
// @Description Create a new note for user
// @Tags Notes
// @Accept json
// @Produce json
// @Param request body web.NoteCreateRequest true "Create Note"
// @Success 200 {object} web.WebResponse
// @Router /notes [post]
// @Security BearerAuth
func (nc *NoteController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request web.NoteCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := nc.NoteService.Create(r.Context(), request)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WebResponse)
}

// GetAll godoc
// @Summary Get all notes
// @Description Get all notes for authorized user
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse
// @Router /notes [get]
// @Security BearerAuth
func (nc *NoteController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := nc.NoteService.GetAll(r.Context())

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WebResponse)
}

// GetbyId godoc
// @Summary Id a note
// @Description Id existing note
// @Tags Notes
// @Accept json
// @Produce json
// @Param request body web.NoteUpdateRequest true "Id Note"
// @Success 200 {object} web.WebResponse
// @Router /notes/{id} [get]
// @Security BearerAuth
func (controller *NoteController) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("id")

    uuidId, err := uuid.Parse(id)
    if err != nil {
        helper.WriteUnauthorized(w, "Invalid UUID format")
        return
    }

    noteResponse := controller.NoteService.FindById(r.Context(), uuidId)
    webResponse := web.WebResponse{
        Code:   200,
        Status: "OK",
        Data:   noteResponse,
    }

    helper.WriteToResponseBody(w, webResponse)
}

// Update godoc
// @Summary Update a note
// @Description Update existing note
// @Tags Notes
// @Accept json
// @Produce json
// @Param request body web.NoteUpdateRequest true "Update Note"
// @Success 200 {object} web.WebResponse
// @Router /notes{id} [put]
// @Security BearerAuth
func (nc *NoteController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request web.NoteUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := nc.NoteService.Update(r.Context(), request)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WebResponse)
}

// Delete godoc
// @Summary Delete a note
// @Description Delete note by ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} web.WebResponse
// @Router /notes/{id} [delete]
// @Security BearerAuth
func (nc *NoteController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idParam := ps.ByName("id")
	id, err := uuid.Parse(idParam)
	helper.PanicIfError(err)

	err = nc.NoteService.Delete(r.Context(), id)
	helper.PanicIfError(err)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, WebResponse)
}
