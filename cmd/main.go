package main

import (
	account_factory "bank_server/internal/account/application/factory"
	"bank_server/internal/adapter/web"
	user_factory "bank_server/internal/user/application/factory"
	"context"
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	dbConn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	createUserUseCase := user_factory.CreateUserUseCase(dbConn)
	getUserUseCase := user_factory.GetUserUseCase(dbConn)

	unitOfWork := account_factory.SetupUnitOfWork(ctx,dbConn)
	createAccountUseCase := account_factory.CreateAccountUseCase(dbConn)
	creditValueUseCase := account_factory.CreditValueUseCase(unitOfWork)
	debitValueUseCase := account_factory.DebitValueUseCase(unitOfWork)
	transferUseCase := account_factory.TransferUseCase(unitOfWork)

	app := web.NewApplication(
		createUserUseCase,
		getUserUseCase,
		createAccountUseCase,
		creditValueUseCase,
		debitValueUseCase,
		transferUseCase,
	)

	app.Server()
}
