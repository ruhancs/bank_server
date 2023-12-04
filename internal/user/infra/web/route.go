package web_user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/health"))

	mux.Post("/user", app.CreateUserHandler)
	mux.Get("/user/{username}", app.GetUserHandler)
	mux.Get("/user/{username}", app.GetUserHandler)

	return mux
}
