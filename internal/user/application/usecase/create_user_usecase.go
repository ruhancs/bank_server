package user_usecase

import (
	user_dto "bank_server/internal/user/application/dto"
	"bank_server/internal/user/domain/entity"
	gateway_user "bank_server/internal/user/domain/gateway"
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

type CreateUserUseCase struct {
	Logs           *zap.Logger
	UserRepository gateway_user.UserRepositoryInterface
}

func NewCreateUserUseCase(repository gateway_user.UserRepositoryInterface, logs *zap.Logger) *CreateUserUseCase {
	return &CreateUserUseCase{
		Logs:           logs,
		UserRepository: repository,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, input user_dto.InputCreateUserDto) (*user_dto.OutputCreateUserDto, error) {
	user, err := entity.NewUser(input.UserName, input.Email)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		return nil, errors.New("failed to create user, please try again later")
	}

	t := time.Now().Format(time.RFC3339)
	u.Logs.Info("User Created",
		zap.String("action", "CreateUserUseCase"),
		zap.String("status", "Success"),
		zap.String("service", "bank_server"),
		zap.String("called at time", t),
	)

	return &user_dto.OutputCreateUserDto{
		//ID: user.ID,
		UserName: user.Username,
		Email:    user.Email,
	}, nil
}
