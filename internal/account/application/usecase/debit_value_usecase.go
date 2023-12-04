package account_usecase

import (
	dto_account "bank_server/internal/account/application/dto"
	"bank_server/internal/account/domain/entity"
	gateway_account "bank_server/internal/account/domain/gateway"
	"bank_server/pkg/uow"
	"context"
	"log"
)


type DebitValueUseCase struct {
	UnitOfWork uow.UowInterface
}

func NewDebitValueUseCase(unitOfWork uow.UowInterface) *DebitValueUseCase{
	return &DebitValueUseCase{
		UnitOfWork: unitOfWork,
	}
}

func(u *DebitValueUseCase) Execute(ctx context.Context,input dto_account.InputDebitValueUseCase) (*dto_account.OutputDebitValueUseCase,error) {
	err := u.UnitOfWork.Do(ctx, func(uow *uow.Uow) error {
		accountRepository := u.getAccountRepository(ctx)
		account,err := accountRepository.GetToUpdate(ctx,input.AccountID)
		if err != nil {
			return err
		}
		
		err = account.DebitBalance(input.Value)
		if err != nil {
			return err
		}
		err = accountRepository.UpdateBalance(ctx,account.ID,account.Balance)
	
		entry,err := account_entity.NewEntry(input.AccountID,account_entity.DEBIT,input.Value)
		if err != nil {
			return err
		}
		entryRepository := u.getEntryRepository(ctx)
		err = entryRepository.Create(ctx,entry)
		if err != nil {
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


