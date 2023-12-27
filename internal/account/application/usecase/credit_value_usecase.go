package account_usecase

import (
	dto_account "bank_server/internal/account/application/dto"
	account_entity "bank_server/internal/account/domain/entity"
	gateway_account "bank_server/internal/account/domain/gateway"
	email "bank_server/internal/adapter/mail"
	"bank_server/pkg/uow"
	"context"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
)

type CreditValueUseCase struct {
	Logger     *zap.Logger
	UnitOfWork uow.UowInterface
	SesMailSender *email.SeSMailSender
}

func NewCreditValueUseCase(unitOfWork uow.UowInterface, logger *zap.Logger, sesMail *email.SeSMailSender) *CreditValueUseCase {
	return &CreditValueUseCase{
		Logger: logger,
		UnitOfWork: unitOfWork,
		SesMailSender: sesMail,
	}
}

func (u *CreditValueUseCase) Execute(ctx context.Context, input dto_account.InputCreditValueUseCase) (*dto_account.OutputCreditValueUseCase, error) {
	err := u.UnitOfWork.Do(ctx, func(uow *uow.Uow) error {
		accountRepository := u.getAccountRepository(ctx)
		account, err := accountRepository.GetToUpdate(ctx, input.AccountID)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to get account to update balance",
				zap.String("action", "CreditValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}

		account.CreditBalance(input.Value)
		err = accountRepository.UpdateBalance(ctx, account.ID, account.Balance)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to update account balance",
				zap.String("action", "CreditValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		
		entry, err := account_entity.NewEntry(input.AccountID, account_entity.CREDIT, input.Value)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to create entry entity",
				zap.String("action", "CreditValueUseCase.Execute"),
				zap.String("status", "Error"),
				zap.String("service", "bank_server"),
				zap.String("called at time", t),
				zap.String("ERROR", err.Error()),
			)
			return err
		}
		
		entryRepository := u.getEntryRepository(ctx)
		err = entryRepository.Create(ctx, entry)
		if err != nil {
			t := time.Now().Format(time.RFC3339)
			u.Logger.Error("Error to save entry on db",
				zap.String("action", "CreditValueUseCase.Execute"),
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
		return nil, err
	}

	output := &dto_account.OutputCreditValueUseCase{
		Status: "success credited",
		Value:  input.Value,
	}

	//enviar email de confirmacao
	u.SesMailSender.Wait.Add(1)
	go u.SesMailSender.SendInvoiceMail(os.Getenv("EMAIL_TO_TESTE"), "Teste Body")

	return output, nil
}

func (u *CreditValueUseCase) getAccountRepository(ctx context.Context) gateway_account.AccountRepositoryInterface {
	repo, err := u.UnitOfWork.GetRepository(ctx, "AccountRepository")
	if err != nil {
		log.Println("failed to get account repository")
		panic(err)
	}
	return repo.(gateway_account.AccountRepositoryInterface)
}

func (u *CreditValueUseCase) getEntryRepository(ctx context.Context) gateway_account.EntryrepositoryInterface {
	repo, err := u.UnitOfWork.GetRepository(ctx, "EntryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(gateway_account.EntryrepositoryInterface)
}
