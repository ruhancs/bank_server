package gateway_user

import (
	"bank_server/internal/user/domain/entity"
	"context"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context,user *entity.User) error
	GetByUserName(ctx context.Context, username string) (*entity.User,error)
}