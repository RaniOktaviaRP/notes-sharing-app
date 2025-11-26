package main

import (
	"fmt"
	"net/http"

	"notes-app/backend/app"
	"notes-app/backend/controller"
	"notes-app/backend/repository"
	"notes-app/backend/service"

	"github.com/joho/godotenv"
	"github.com/go-playground/validator/v10"

	_ "notes-app/backend/docs"
)

// @title Notes API
// @version 1.0
// @description API for Note Application
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    _ = godotenv.Load(".env") 

	db := app.NewDB()
	defer db.Close()

	validate := validator.New()

	userRepository := repository.NewUserRepository()
	noteRepository := repository.NewNoteRepositoryImpl(db)

	userService := service.NewUserServiceImpl(userRepository, db)
	noteService := service.NewNoteServiceImpl(noteRepository, db, validate)

	userController := controller.NewUserController(userService)
	noteController := controller.NewNoteController(noteService)

	router := app.NewRouter(userController, noteController)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
