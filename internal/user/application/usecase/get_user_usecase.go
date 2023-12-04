package user_usecase

import (
	user_dto "bank_server/internal/user/application/dto"
	gateway_user "bank_server/internal/user/domain/gateway"
	"context"
)

type GetUserUseCase struct {
	UserRepository gateway_user.UserRepositoryInterface
}

func NewGetUserUseCase(repo gateway_user.UserRepositoryInterface) *GetUserUseCase{
	return &GetUserUseCase{
		UserRepository: repo,
	}
}

func(u *GetUserUseCase) Execute(ctx context.Context,username string) (*user_dto.OutputGetUserDto,error) {
	user,err := u.UserRepository.GetByUserName(ctx,username)
	if err != nil {
		return nil,err
	}

	output := &user_dto.OutputGetUserDto{
		UserName: user.Username,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}

	return output,nil
}