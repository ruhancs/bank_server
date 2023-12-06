package gateway_account

import (
	account_entity "bank_server/internal/account/domain/entity"
	"context"
)

type EntryrepositoryInterface interface {
	Create(ctx context.Context, entry *account_entity.Entry) error
	BulkCreate(ctx context.Context, fromEntry *account_entity.Entry, toEntry *account_entity.Entry) error
	Get(ctx context.Context, id string) (*account_entity.Entry, error)
	List(ctx context.Context, accountID string, perPage int, page int) ([]account_entity.Entry, error)
	Delete(ctx context.Context, id string) error
}
