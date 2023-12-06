package account_usecase

import (
	dto_account "bank_server/internal/account/application/dto"
	account_entity "bank_server/internal/account/domain/entity"
	gateway_account "bank_server/internal/account/domain/gateway"
	"bank_server/pkg/uow"
	"context"
	"errors"
	"log"
	"time"

	"go.uber.org/zap"
)

type TransferUseCase struct {
	Logger     *zap.Logger
	UnitOfWork uow.UowInterface
}

func NewTransferUseCase(unitOfWork uow.UowInterface, logger *zap.Logger) *TransferUseCase {
	return &TransferUseCase{
		Logger:     logger,
		UnitOfWork: unitOfWork,
	}
}

func (u *TransferUseCase) Execute(ctx context.Context, input dto_account.InputTransferUseCase) (dto_account.OutputTransferUseCase, error) {
	err := u.UnitOfWork.Do(ctx, func(uow *uow.Uow) error {
		accountRepo := u.getAccountRepository(ctx)
		entryRepo := u.getEntryRepository(ctx)
		transferRepo := u.getTransferRepository(ctx)

		fromAccount, err := accountRepo.GetToUpdate(ctx, input.FromAccountID)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to get from account transfer to update",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}

		toAccount, err := accountRepo.GetToUpdate(ctx, input.ToAccountID)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to get to account transfer to update",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}

		err = fromAccount.DebitBalance(input.Value)
		if err != nil {
			return err
		}
		toAccount.CreditBalance(input.Value)

		//refactor to bulkupdate
		err = accountRepo.UpdateBalance(ctx, input.FromAccountID, fromAccount.Balance)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to update balance in from account transfer",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		err = accountRepo.UpdateBalance(ctx, input.ToAccountID, toAccount.Balance)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to update balance in to account transfer",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}

		entryFromAccount, err := account_entity.NewEntry(fromAccount.ID, account_entity.DEBIT, input.Value)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to create entry entity for from account transfer",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		entryToAccount, err := account_entity.NewEntry(toAccount.ID, account_entity.CREDIT, input.Value)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to create entry entity for to account transfer",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}

		err = entryRepo.BulkCreate(ctx, entryFromAccount, entryToAccount)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to bulk create entries on db",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}

		transfer, err := account_entity.NewTransfer(fromAccount.ID, toAccount.ID, input.Value)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to create transfer entity",
				zap.String("action", "TransferUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		err = transferRepo.Create(ctx, transfer)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to save transfer on db",
				zap.String("action", "TransferUseCase.Execute"),
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
		return dto_account.OutputTransferUseCase{
			FromAccountID: input.FromAccountID,
			ToAccountID:   input.ToAccountID,
			Value:         input.Value,
			Status:        "failed to transfer",
		}, err
	}

	return dto_account.OutputTransferUseCase{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Value:         input.Value,
		Status:        "success",
	}, nil
}

func (u *TransferUseCase) checkFromAccountBalance(account *account_entity.Account, value int) error {
	if account.Balance < value {
		return errors.New("insuficient balance to transfer")
	}
	return nil
}

func (u *TransferUseCase) getAccountRepository(ctx context.Context) gateway_account.AccountRepositoryInterface {
	repo, err := u.UnitOfWork.GetRepository(ctx, "AccountRepository")
	if err != nil {
		log.Println("failed to get account repository")
		panic(err)
	}
	return repo.(gateway_account.AccountRepositoryInterface)
}

func (u *TransferUseCase) getEntryRepository(ctx context.Context) gateway_account.EntryrepositoryInterface {
	repo, err := u.UnitOfWork.GetRepository(ctx, "EntryRepository")
	if err != nil {
		log.Println("failed to get entry repository")
		panic(err)
	}
	return repo.(gateway_account.EntryrepositoryInterface)
}

func (u *TransferUseCase) getTransferRepository(ctx context.Context) gateway_account.TransferRepositoryInterface {
	repo, err := u.UnitOfWork.GetRepository(ctx, "TransferRepository")
	if err != nil {
		log.Println("failed to get entry repository")
		panic(err)
	}
	return repo.(gateway_account.TransferRepositoryInterface)
}
