package user_usecase

import (
	user_dto "bank_server/internal/user/application/dto"
	gateway_user "bank_server/internal/user/domain/gateway"
	"context"
	"time"

	"go.uber.org/zap"
)

type GetUserUseCase struct {
	Logs           *zap.Logger
	UserRepository gateway_user.UserRepositoryInterface
}

func NewGetUserUseCase(repo gateway_user.UserRepositoryInterface, logger *zap.Logger) *GetUserUseCase {
	return &GetUserUseCase{
		Logs:           logger,
		UserRepository: repo,
	}
}

func (u *GetUserUseCase) Execute(ctx context.Context, username string) (*user_dto.OutputGetUserDto, error) {
	user, err := u.UserRepository.GetByUserName(ctx, username)
	if err != nil {
		return nil, err
	}

	t := time.Now().Format(time.RFC3339)
	u.Logs.Info("Get User",
		zap.String("action", "GetUserUseCase.Execute"),
		zap.String("status", "Success"),
		zap.String("service", "bank_server"),
		zap.String("called at time", t),
	)

	output := &user_dto.OutputGetUserDto{
		UserName:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return output, nil
}
