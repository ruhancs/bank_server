package web_user

import (
	user_usecase "bank_server/internal/user/application/usecase"
	"log"
	"net/http"
	"time"
)

type Application struct {
	CreateUserUseCase *user_usecase.CreateUserUseCase
	GetUserUseCase    *user_usecase.GetUserUseCase
}

func NewApplication(createUserUseCase *user_usecase.CreateUserUseCase,getUserUseCase *user_usecase.GetUserUseCase) *Application{
	return &Application{
		CreateUserUseCase: createUserUseCase,
		GetUserUseCase: getUserUseCase,
	}
}

func (app *Application) Server() error {
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Println("Runing server on port 8000...")
	return srv.ListenAndServe()
}
