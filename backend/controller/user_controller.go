package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"notes-app/backend/model/web"
	"notes-app/backend/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

// Register handles user registration requests
// @Summary Register a new user
// @Description Register a new user with email, name, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param request body web.UserRegisterRequest true "User Register Request"
// @Success 201 {object} web.WebResponse{data=web.UserResponse} "User registered successfully"
// @Failure 400 {object} web.WebResponse{message=string} "Bad Request"
// @Failure 500 {object} web.WebResponse{message=string} "Internal Server Error"
// @Router /users [post]
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request web.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := uc.UserService.Register(r.Context(), request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	WebResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data:   res,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WebResponse)
}	

// Login
// @Summary Login a user
// @Description Login a user with email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param request body web.UserLoginRequest true "User Login Request"
// @Success 200 {object} web.WebResponse{data=string} "User logged in successfully"
// @Failure 400 {object} web.WebResponse{message=string} "Bad Request"
// @Failure 401 {object} web.WebResponse{message=string} "Unauthorized"
// @Failure 500 {object} web.WebResponse{message=string} "Internal Server Error"
// @Router /login [post]
func (c *UserController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req web.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	token, err := c.UserService.Login(r.Context(), req)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	WebResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WebResponse)
}