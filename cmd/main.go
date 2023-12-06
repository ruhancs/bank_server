package main

import (
	account_factory "bank_server/internal/account/application/factory"
	"bank_server/internal/adapter/web"
	user_factory "bank_server/internal/user/application/factory"
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// @title           Bank Server
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Ruhan CS
// @contact.url    ruhancorreasoares@gmail.com
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

func main() {
	
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	
	ctx := context.Background()
	//dbConn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE_DOCKER"))
	dbConn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		panic(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Println("cannot connect to db")
		panic(err)
	}
	defer dbConn.Close()

	logger,_ := zap.NewProduction()
	defer logger.Sync()

	createUserUseCase := user_factory.CreateUserUseCase(dbConn,logger)
	getUserUseCase := user_factory.GetUserUseCase(dbConn,logger)

	unitOfWork := account_factory.SetupUnitOfWork(ctx,dbConn)
	createAccountUseCase := account_factory.CreateAccountUseCase(dbConn)
	creditValueUseCase := account_factory.CreditValueUseCase(unitOfWork,logger)
	debitValueUseCase := account_factory.DebitValueUseCase(unitOfWork,logger)
	transferUseCase := account_factory.TransferUseCase(unitOfWork,logger)

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
