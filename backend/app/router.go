package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"

	"notes-app/backend/controller"
	"notes-app/backend/middleware"
)

func WrapHandlerWithMiddleware(handler http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handler.ServeHTTP(w, r)
	}
}

func WrapHandlerWithJWT(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, ps)
		})
		middleware.JWTAuth(h).ServeHTTP(w, r)
	}
}

func NewRouter(userController controller.UserController) *httprouter.Router {
    router := httprouter.New()

    router.POST("/users", WrapHandlerWithMiddleware(
        middleware.CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            userController.Register(w, r, httprouter.Params{})
        })),
    ))

    router.POST("/login", WrapHandlerWithMiddleware(
        middleware.CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            userController.Login(w, r, httprouter.Params{})
        })),
    ))

	router.POST("/notes", WrapHandlerWithJWT(
        middleware.CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            noteController.Create(w, r, httprouter.Params{})
        })),
    ))

	router.PUT("/notes/:id", WrapHandlerWithJWT(
        middleware.CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            noteController.Update(w, r, httprouter.Params{})
        })),
    ))

	router.DELETE("/notes/:id", WrapHandlerWithJWT(
        middleware.CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            noteController.Delete(w, r, httprouter.Params{})
        })),
    ))

	router.GET("/swagger/*any", WrapHandlerWithMiddleware(
		middleware.CORS(httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		)),
	))

	return router
}
