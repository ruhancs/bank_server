package account_factory

import (
	account_usecase "bank_server/internal/account/application/usecase"
	account_repository "bank_server/internal/account/infra/repository"
	email "bank_server/internal/adapter/mail"
	user_repository "bank_server/internal/user/infra/repository"
	"bank_server/pkg/uow"
	"bank_server/sql/db"
	"context"
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func SetupUnitOfWork(ctx context.Context,database *sql.DB) *uow.Uow {
	uow := uow.NewUow(ctx,database)
	uow.Register("AccountRepository", func (tx *sql.Tx) interface{}  {
		repo := account_repository.NewAccountRepository(database)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("EntryRepository", func (tx *sql.Tx) interface{}  {
		repo := account_repository.NewEntryRepository(database)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("TransferRepository", func (tx *sql.Tx) interface{}  {
		repo := account_repository.NewTransferRepository(database)
		repo.Queries = db.New(tx)
		return repo
	})
	return uow
}

func CreateAccountUseCase(db *sql.DB) *account_usecase.CreateAccountUseCase {
	accountRepository := account_repository.NewAccountRepository(db)
	userRepository := user_repository.NewUserRepository(db)
	usecase := account_usecase.NewCreateAccountUseCase(accountRepository,userRepository)
	return usecase
}

func CreditValueUseCase(unitOfWork uow.UowInterface, logger *zap.Logger, sesMail *email.SeSMailSender) *account_usecase.CreditValueUseCase {
	usecase := account_usecase.NewCreditValueUseCase(unitOfWork,logger,sesMail)
	return usecase
}

func DebitValueUseCase(unitOfWork uow.UowInterface, logger *zap.Logger) *account_usecase.DebitValueUseCase {
	usecase := account_usecase.NewDebitValueUseCase(unitOfWork,logger)
	return usecase
}

func TransferUseCase(
	transferErrorsGetFromAccount prometheus.Counter,
	transferErrorsGetToAccount prometheus.Counter,
	transferErrorsUpdateFromAccountBalance prometheus.Counter,
	transferErrorsUpdateToAccountBalance prometheus.Counter,
	transferErrorsCreateEntries prometheus.Counter,
	transferErrorsCreateTransfer prometheus.Counter,
	unitOfWork uow.UowInterface,
	logger *zap.Logger,
) *account_usecase.TransferUseCase {
	usecase := account_usecase.NewTransferUseCase(
		transferErrorsGetFromAccount,
		transferErrorsGetToAccount,
		transferErrorsUpdateFromAccountBalance,
		transferErrorsUpdateToAccountBalance,
		transferErrorsCreateEntries,
		transferErrorsCreateTransfer,
		unitOfWork,
		logger,
	)
	return usecase
}