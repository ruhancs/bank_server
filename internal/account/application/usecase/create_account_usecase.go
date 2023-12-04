package account_usecase

import (
	dto_account "bank_server/internal/account/application/dto"
	"bank_server/internal/account/domain/entity"
	gateway_account "bank_server/internal/account/domain/gateway"
	gateway_user "bank_server/internal/user/domain/gateway"
	"context"
	"errors"
)

type CreateAccountUseCase struct {
	AccountRepository gateway_account.AccountRepositoryInterface
	UserRepository    gateway_user.UserRepositoryInterface
}

func NewCreateAccountUseCase(accountRepo gateway_account.AccountRepositoryInterface, userRepo gateway_user.UserRepositoryInterface) *CreateAccountUseCase{
	return &CreateAccountUseCase{
		AccountRepository: accountRepo,
		UserRepository: userRepo,
	}
}

func(u *CreateAccountUseCase) Execute(ctx context.Context,input dto_account.InputCreateAccountDto) (*dto_account.OutputCreateAccountDto,error) {
	_,err := u.UserRepository.GetByUserName(ctx,input.Owner)
	if err != nil {
		return nil, errors.New("user not registered")
	}

	account,err := account_entity.NewAccount(input.Owner)
	if err != nil {
		return nil, errors.New("failed to create account, please check the input data")
	}
	
	err = u.AccountRepository.Create(ctx,account)
	if err != nil {
		return nil, errors.New("failed to create account, please try again latter")
	}

	return &dto_account.OutputCreateAccountDto{
		ID: account.ID,
		Owner: account.Owner,
		Balance: account.Balance,
		CreateAt: account.CreatedAt,
	}, nil
}
