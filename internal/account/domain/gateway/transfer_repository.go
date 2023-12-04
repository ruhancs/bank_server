package gateway_account

import (
	account_entity "bank_server/internal/account/domain/entity"
	"context"
)

type TransferRepositoryInterface interface {
	Create(ctx context.Context, transfer *account_entity.Transfer) error
	Get(ctx context.Context, id string) (*account_entity.Transfer, error)
	List(ctx context.Context, perPage, page int) ([]account_entity.Transfer, error)
	//ListByToAccount(fromAccountID,toAccountID string) ([]*account_entity.Transfer,error)
	Delete(ctx context.Context, id string) error
}
