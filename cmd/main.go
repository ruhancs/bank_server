package main

import (
	account_factory "bank_server/internal/account/application/factory"
	email "bank_server/internal/adapter/mail"
	"bank_server/internal/adapter/web"
	user_factory "bank_server/internal/user/application/factory"
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	transferErrorsGetFromAccount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transfer_errors_get_from_account",
		Help: "Error in method in repository GetToUpdate with from account",
	})
	transferErrorsGetToAccount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transfer_errors_get_to_account",
		Help: "Error in method in repository GetToUpdate with to account",
	})
	transferErrorsUpdateFromAccountBalance = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transfer_errors_update_from_account_balance",
		Help: "Error in method in repository UpdateBalance with from account",
	})
	transferErrorsUpdateToAccountBalance = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transfer_errors_update_to_account_balance",
		Help: "Error in method in repository UpdateBalance with to account",
	})
	transferErrorsCreateEntries = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transfer_errors_create_entries",
		Help: "Error in method from entry repository BulkCreate",
	})
	transferErrorsCreateTransfer = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transfer_errors_create_transfer",
		Help: "Error in method from transfer repository Create",
	})
	transferHandlerErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transfer_handler_errors",
		Help: "Errors in handler transfer",
	})
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .Env file")
	}
	prometheus.MustRegister(transferErrorsGetFromAccount)
	prometheus.MustRegister(transferErrorsGetToAccount)
	prometheus.MustRegister(transferErrorsUpdateFromAccountBalance)
	prometheus.MustRegister(transferErrorsUpdateToAccountBalance)
	prometheus.MustRegister(transferErrorsCreateEntries)
	prometheus.MustRegister(transferErrorsCreateTransfer)
	prometheus.MustRegister(transferHandlerErrors)
}

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

	mailErrChan := make(chan error)
	waitGroup := sync.WaitGroup{}
	sesSession := email.CreateSession(os.Getenv("AWS_REGION"),os.Getenv("PK"),os.Getenv("SK"))
	sesMail := email.NewSesMailSender(sesSession,&waitGroup,mailErrChan)
	go sesMail.ListenForMail()

	createUserUseCase := user_factory.CreateUserUseCase(dbConn,logger)
	getUserUseCase := user_factory.GetUserUseCase(dbConn,logger)

	unitOfWork := account_factory.SetupUnitOfWork(ctx,dbConn)
	createAccountUseCase := account_factory.CreateAccountUseCase(dbConn)
	creditValueUseCase := account_factory.CreditValueUseCase(unitOfWork,logger,sesMail)
	debitValueUseCase := account_factory.DebitValueUseCase(unitOfWork,logger)
	transferUseCase := account_factory.TransferUseCase(
		transferErrorsGetFromAccount,
		transferErrorsGetToAccount,
		transferErrorsUpdateFromAccountBalance,
		transferErrorsUpdateToAccountBalance,
		transferErrorsCreateEntries,
		transferErrorsCreateTransfer,
		unitOfWork,
		logger,
	)

	app := web.NewApplication(
		transferHandlerErrors,
		createUserUseCase,
		getUserUseCase,
		createAccountUseCase,
		creditValueUseCase,
		debitValueUseCase,
		transferUseCase,
	)

	http.Handle("/metrics", promhttp.Handler())
	go func ()  {
		log.Println("Running metrics")
		http.ListenAndServe(":8080", nil)
	}()

	app.Server()
}
