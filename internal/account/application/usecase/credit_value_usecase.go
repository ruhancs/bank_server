package account_usecase

import (
	dto_account "bank_server/internal/account/application/dto"
	"bank_server/internal/account/domain/entity"
	gateway_account "bank_server/internal/account/domain/gateway"
	"bank_server/pkg/uow"
	"context"
	"log"
)


type CreditValueUseCase struct {
	UnitOfWork uow.UowInterface
}

func NewCreditValueUseCase(unitOfWork uow.UowInterface) *CreditValueUseCase{
	return &CreditValueUseCase{
		UnitOfWork: unitOfWork,
	}
}

func(u *CreditValueUseCase) Execute(ctx context.Context,input dto_account.InputCreditValueUseCase) (*dto_account.OutputCreditValueUseCase,error) {
	err := u.UnitOfWork.Do(ctx, func(uow *uow.Uow) error {
		accountRepository := u.getAccountRepository(ctx)
		account,err := accountRepository.GetToUpdate(ctx,input.AccountID)
		if err != nil {
			return err
		}
	
		account.CreditBalance(input.Value)
		err = accountRepository.UpdateBalance(ctx,account.ID,account.Balance)
		if err != nil {
			return err
		}
	
		entry,err := account_entity.NewEntry(input.AccountID,account_entity.CREDIT,input.Value)
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

	output := &dto_account.OutputCreditValueUseCase{
		Status: "success credited",
		Value: input.Value,
	}

	return output,nil
}

func(u *CreditValueUseCase) getAccountRepository(ctx context.Context) gateway_account.AccountRepositoryInterface {
	log.Println("GET ACCOUNT REPO")
	repo,err := u.UnitOfWork.GetRepository(ctx, "AccountRepository")
	if err != nil {
		log.Println("failed to get account repository")
		panic(err)
	}
	log.Println("GET ACCOUNT REPO FINISH")
	return repo.(gateway_account.AccountRepositoryInterface)
}

func(u *CreditValueUseCase) getEntryRepository(ctx context.Context) gateway_account.EntryrepositoryInterface {
	log.Println("GET ENTRY REPO")
	repo,err := u.UnitOfWork.GetRepository(ctx, "EntryRepository")
	if err != nil {
		log.Println("failed to get entry repository_______________")
		panic(err)
	}
	log.Println("GET ENTRY REPO FINISH")
	return repo.(gateway_account.EntryrepositoryInterface)
}


