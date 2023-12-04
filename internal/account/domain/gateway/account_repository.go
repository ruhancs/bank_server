package gateway_account

import (
	"bank_server/internal/account/domain/entity"
	"context"
)

type AccountRepositoryInterface interface {
	Create(ctx context.Context, account *account_entity.Account) error
	Get(ctx context.Context, id string) (*account_entity.Account, error)
	GetToUpdate(ctx context.Context, id string) (*account_entity.Account, error)
	UpdateBalance(ctx context.Context, id string, balance int) error
	Delete(ctx context.Context, accountID string) error
}
