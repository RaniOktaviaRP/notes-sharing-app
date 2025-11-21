package app 

import (
	"net/http"
	"notes-app/backend/controller"
	httpSwagger "github.com/swaggo/http-swagger"

)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/users/register", userController.Register)
	router.GET("/swagger/*any", httpSwagger.WrapHandler)
	return router
}