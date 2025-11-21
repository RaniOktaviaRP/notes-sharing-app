package controller

import (
	"notes-app/backend/model/web"
	"notes-app/backend/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
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
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var request web.UserRegisterRequest
	err := utils.ReadRequestBody(r, &request)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	res := c.UserService.Register(r.Context(), request)
	WebResponse := web.WebResponse{
		Message: "User registered successfully",
		Data:    res,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WebResponse)
}	