package account_usecase

import (
	dto_account "bank_server/internal/account/application/dto"
	"bank_server/internal/account/domain/entity"
	gateway_account "bank_server/internal/account/domain/gateway"
	"bank_server/pkg/uow"
	"context"
	"errors"
	"log"
)

type TransferUseCase struct {
	UnitOfWork uow.UowInterface
}

func NewTransferUseCase(unitOfWork uow.UowInterface) *TransferUseCase {
	return &TransferUseCase{
		UnitOfWork: unitOfWork,
	}
}

func (u *TransferUseCase) Execute(ctx context.Context, input dto_account.InputTransferUseCase) (dto_account.OutputTransferUseCase, error) {
	err := u.UnitOfWork.Do(ctx, func(uow *uow.Uow) error {
		accountRepo := u.getAccountRepository(ctx)
		entryRepo := u.getEntryRepository(ctx)
		transferRepo := u.getTransferRepository(ctx)

		fromAccount, err := accountRepo.GetToUpdate(ctx,input.FromAccountID)
		if err != nil {
			return err
		}

		toAccount, err := accountRepo.GetToUpdate(ctx,input.ToAccountID)
		if err != nil {
			return err
		}
		
		err = fromAccount.DebitBalance(input.Value)
		if err != nil {
			return err
		}
		toAccount.CreditBalance(input.Value)

		//refactor to bulkupdate
		err = accountRepo.UpdateBalance(ctx,input.FromAccountID, fromAccount.Balance)
		if err != nil {
			return err
		}
		err = accountRepo.UpdateBalance(ctx,input.ToAccountID, toAccount.Balance)
		if err != nil {
			return err
		}

		entryFromAccount, err := account_entity.NewEntry(fromAccount.ID, account_entity.DEBIT, input.Value)
		if err != nil {
			return err
		}
		entryToAccount, err := account_entity.NewEntry(fromAccount.ID, account_entity.CREDIT, input.Value)
		if err != nil {
			return err
		}
		//refactor to bulkcreate
		err = entryRepo.Create(ctx,entryFromAccount)
		if err != nil {
			return err
		}
		err = entryRepo.Create(ctx,entryToAccount)
		if err != nil {
			return err
		}

		transfer, err := account_entity.NewTransfer(fromAccount.ID, toAccount.ID, input.Value)
		if err != nil {
			return err
		}
		err = transferRepo.Create(ctx,transfer)
		if err != nil {
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

func(u *TransferUseCase) checkFromAccountBalance(account *account_entity.Account, value int) error {
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
