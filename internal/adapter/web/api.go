package web

import (
	account_usecase "bank_server/internal/account/application/usecase"
	user_usecase "bank_server/internal/user/application/usecase"
	"log"
	"net/http"
	"time"
)

type Application struct {
	CreateUserUseCase    *user_usecase.CreateUserUseCase
	GetUserUseCase       *user_usecase.GetUserUseCase
	CreateAccountUseCase *account_usecase.CreateAccountUseCase
	CreditValueUseCase   *account_usecase.CreditValueUseCase
	DebitValueUseCase    *account_usecase.DebitValueUseCase
	TransferUseCase      *account_usecase.TransferUseCase
}

func NewApplication(
	createUserUseCase *user_usecase.CreateUserUseCase,
	getUserUseCase *user_usecase.GetUserUseCase,
	createAccountUseCase *account_usecase.CreateAccountUseCase,
	creditValuetUseCase *account_usecase.CreditValueUseCase,
	debitValueUseCase *account_usecase.DebitValueUseCase,
	transferUseCase *account_usecase.TransferUseCase,
) *Application {
	return &Application{
		CreateUserUseCase:    createUserUseCase,
		GetUserUseCase:       getUserUseCase,
		CreateAccountUseCase: createAccountUseCase,
		CreditValueUseCase:   creditValuetUseCase,
		DebitValueUseCase:    debitValueUseCase,
		TransferUseCase:      transferUseCase,
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
