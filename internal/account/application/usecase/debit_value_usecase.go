package account_usecase

import (
	dto_account "bank_server/internal/account/application/dto"
	"bank_server/internal/account/domain/entity"
	gateway_account "bank_server/internal/account/domain/gateway"
	"bank_server/pkg/uow"
	"context"
	"log"
	"time"

	"go.uber.org/zap"
)


type DebitValueUseCase struct {
	Logger     *zap.Logger
	UnitOfWork uow.UowInterface
}

func NewDebitValueUseCase(unitOfWork uow.UowInterface, logger *zap.Logger) *DebitValueUseCase{
	return &DebitValueUseCase{
		Logger: logger,
		UnitOfWork: unitOfWork,
	}
}

func(u *DebitValueUseCase) Execute(ctx context.Context,input dto_account.InputDebitValueUseCase) (*dto_account.OutputDebitValueUseCase,error) {
	err := u.UnitOfWork.Do(ctx, func(uow *uow.Uow) error {
		accountRepository := u.getAccountRepository(ctx)
		account,err := accountRepository.GetToUpdate(ctx,input.AccountID)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to get account to update",
				zap.String("action", "DebitValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		
		err = account.DebitBalance(input.Value)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to debit value in account",
				zap.String("action", "DebitValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		err = accountRepository.UpdateBalance(ctx,account.ID,account.Balance)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to update balance of the account in db",
				zap.String("action", "DebitValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
	
		entry,err := account_entity.NewEntry(input.AccountID,account_entity.DEBIT,input.Value)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to create entry entity",
				zap.String("action", "DebitValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		entryRepository := u.getEntryRepository(ctx)
		err = entryRepository.Create(ctx,entry)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to save entry on db",
				zap.String("action", "DebitValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		return nil
	})

	if err != nil {
		return nil,err
	}

	output := &dto_account.OutputDebitValueUseCase{
		Status: "success debited",
		Value: input.Value,
	}

	return output,nil
}

func(u *DebitValueUseCase) getAccountRepository(ctx context.Context) gateway_account.AccountRepositoryInterface {
	repo,err := u.UnitOfWork.GetRepository(ctx, "AccountRepository")
	if err != nil {
		log.Println("failed to get account repository")
		panic(err)
	}
	return repo.(gateway_account.AccountRepositoryInterface)
}

func(u *DebitValueUseCase) getEntryRepository(ctx context.Context) gateway_account.EntryrepositoryInterface {
	repo,err := u.UnitOfWork.GetRepository(ctx, "EntryRepository")
	if err != nil {
		log.Println("failed to get entry repository")
		panic(err)
	}
	return repo.(gateway_account.EntryrepositoryInterface)
}


