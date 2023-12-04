package user_usecase

import (
	"bank_server/internal/user/application/dto"
	"bank_server/internal/user/domain/entity"
	gateway_user "bank_server/internal/user/domain/gateway"
	"context"
	"errors"
)

type CreateUserUseCase struct {
	UserRepository gateway_user.UserRepositoryInterface
}

func NewCreateUserUseCase(repository gateway_user.UserRepositoryInterface) *CreateUserUseCase{
	return &CreateUserUseCase{
		UserRepository: repository,
	}
}

func(u *CreateUserUseCase) Execute(ctx context.Context,input user_dto.InputCreateUserDto) (*user_dto.OutputCreateUserDto,error) {
	user,err := entity.NewUser(input.UserName,input.Email)
	if err != nil {
		return nil,err
	}

	err = u.UserRepository.Create(ctx,user)
	if err != nil {
		return nil,errors.New("failed to create user, please try again later")
	}

	return &user_dto.OutputCreateUserDto{
		//ID: user.ID,
		UserName: user.Username,
		Email: user.Email,
	},nil
}