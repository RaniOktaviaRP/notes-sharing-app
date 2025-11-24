package main 

import (
	"notes-app/backend/app"
	"notes-app/backend/controller"
	"notes-app/backend/service"
	"notes-app/backend/repository"
	"net/http"
	"fmt"
)

func main() {
	userService := service.NewUserService()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)


	router := app.NewRouter(userController)

	http.ListenAndServe(":8080", router)
	fmt.Println("Server is running on port 8080")
} 
