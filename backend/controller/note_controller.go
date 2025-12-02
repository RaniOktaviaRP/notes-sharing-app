package controller

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/julienschmidt/httprouter"

    "notes-app/backend/helper"
    "notes-app/backend/model/web"
    "notes-app/backend/service"
)

type NoteController struct {
    NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
    return NoteController{NoteService: noteService}
}

// Create godoc
// @Summary Create a new note
// @Description Create a new note for user
// @Tags Notes
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Title"
// @Param content formData string true "Content"
// @Param image formData file false "Image file"
// @Success 200 {object} web.WebResponse
// @Router /notes [post]
// @Security BearerAuth
func (nc *NoteController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		helper.WriteBadRequest(w, "Invalid form-data")
		return
	}

	_, fileHeader, _ := r.FormFile("image")

	request := web.NoteCreateRequest{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
		Image:   fileHeader,
	}

	res, err := nc.NoteService.Create(r.Context(), request)
	if err != nil {
		helper.WriteBadRequest(w, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// GetAll godoc
// @Summary Get all notes
// @Description Get all notes for authorized user
// @Tags Notes
// @Produce json
// @Success 200 {object} web.WebResponse
// @Router /notes [get]
// @Security BearerAuth
func (nc *NoteController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	res, err := nc.NoteService.GetAll(r.Context())
	if err != nil {
		helper.WriteBadRequest(w, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// FindById godoc
// @Summary Get note by ID
// @Description Get a note by its ID
// @Tags Notes
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} web.WebResponse
// @Router /notes/{id} [get]
// @Security BearerAuth
func (nc *NoteController) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    id := ps.ByName("id")
    uuidId, err := uuid.Parse(id)
    if err != nil {
        helper.WriteBadRequest(w, "Invalid UUID")
        return
    }

    noteResponse, err := nc.NoteService.FindById(r.Context(), uuidId)
    if err != nil {
        helper.WriteNotFound(w, "Note not found")
        return
    }

    helper.WriteToResponseBody(w, web.WebResponse{
        Code:   200,
        Status: "OK",
        Data:   noteResponse,
    })
}

// Update godoc
// @Summary Update a note
// @Description Update an existing note
// @Tags Notes
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Note ID"
// @Param title formData string false "Title"
// @Param content formData string false "Content"
// @Param image formData file false "Image file"
// @Success 200 {object} web.WebResponse
// @Router /notes/{id} [put]
// @Security BearerAuth
func (nc *NoteController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")
	uuidId, err := uuid.Parse(id)
	if err != nil {
		helper.WriteBadRequest(w, "Invalid UUID")
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		helper.WriteBadRequest(w, "Invalid form-data")
		return
	}

	_, fileHeader, _ := r.FormFile("image")

	// pakai pointer utk partial update
	var titlePtr *string
	var contentPtr *string

	titleVal := r.FormValue("title")
	if titleVal != "" {
		titlePtr = &titleVal
	}

	contentVal := r.FormValue("content")
	if contentVal != "" {
		contentPtr = &contentVal
	}

	request := web.NoteUpdateRequest{
		Id:      uuidId,
		Title:   titlePtr,
		Content: contentPtr,
		Image:   fileHeader,
	}

	res, err := nc.NoteService.Update(r.Context(), request)
	if err != nil {
		helper.WriteBadRequest(w, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// Delete godoc
// @Summary Delete a note
// @Description Delete note by ID
// @Tags Notes
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} web.WebResponse
// @Router /notes/{id} [delete]
// @Security BearerAuth
func (nc *NoteController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    idParam := ps.ByName("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        helper.WriteBadRequest(w, "Invalid UUID")
        return
    }

    err = nc.NoteService.Delete(r.Context(), id)
    helper.PanicIfError(err)

    helper.WriteToResponseBody(w, web.WebResponse{
        Code:   200,
        Status: "OK",
    })
}
