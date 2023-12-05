package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	_ "bank_server/docs" //docs do swagger
	httpSwagger "github.com/swaggo/http-swagger" // rota do swagger
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/health"))

	//swagger
	//http://localhost:8000/docs/doc.json/index.html
	//http://localhost:8000/docs/doc.json/index.html#/
	mux.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	mux.Post("/user", app.CreateUserHandler)
	mux.Get("/user/{username}", app.GetUserHandler)

	mux.Post("/account", app.CreateAccountHandler)
	mux.Post("/account/credit", app.CreditValueHandler)
	mux.Post("/account/debit", app.DebitValueHandler)
	mux.Post("/account/transfer", app.TransferHandler)

	return mux
}
