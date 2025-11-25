package main

import (
	"fmt"
	"net/http"

	"notes-app/backend/app"
	"notes-app/backend/controller"
	"notes-app/backend/repository"
	"notes-app/backend/service"
	"github.com/joho/godotenv"
	_ "notes-app/backend/docs"
)

func main() {
	godotenv.Load()

	db := app.NewDB()
	defer db.Close()

	userRepository := repository.NewUserRepository()
	noteRepository := repository.NewNoteRepositoryImpl(db)

	userService := service.NewUserServiceImpl(userRepository, db)
	noteService := service.NewNoteServiceImpl(noteRepository, db)

	userController := controller.NewUserController(userService)
	noteController := controller.NewNoteController(noteService)


	router := app.NewRouter(userController, noteController)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
