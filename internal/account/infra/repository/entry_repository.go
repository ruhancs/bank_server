package account_repository

import (
	account_entity "bank_server/internal/account/domain/entity"
	"bank_server/sql/db"
	"context"
	"database/sql"
	"errors"
)

type EntryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewEntryRepository(database *sql.DB) *EntryRepository {
	return &EntryRepository{
		DB:      database,
		Queries: db.New(database),
	}
}

func (r *EntryRepository) Create(ctx context.Context, entry *account_entity.Entry) error {
	err := r.Queries.CreateEntry(ctx, db.CreateEntryParams{
		ID:              entry.ID,
		AccountID:       entry.AccountID,
		TransactionType: entry.TransactionType,
		Amount:          int64(entry.Amount),
	})
	if err != nil {
		return errors.New("failed to create entry")
	}

	return nil
}

func (r *EntryRepository) BulkCreate(ctx context.Context, fromEntry *account_entity.Entry, toEntry *account_entity.Entry) error {
	err := r.Queries.BulkCreateEntry(ctx,db.BulkCreateEntryParams{
		ID: fromEntry.ID,
		AccountID: fromEntry.AccountID,
		TransactionType: fromEntry.TransactionType,
		Amount: int64(fromEntry.Amount),
		ID_2: toEntry.ID,
		AccountID_2: toEntry.AccountID,
		TransactionType_2: toEntry.TransactionType,
		Amount_2: int64(toEntry.Amount),
	})
	if err != nil {
		return errors.New("failed to create entry")
	}

	return nil
}

func (r *EntryRepository) List(ctx context.Context, accountID string, perPage int, page int) ([]account_entity.Entry, error) {
	offset := (page - 1) * perPage
	entriesModel,err := r.Queries.ListEntry(ctx,db.ListEntryParams{
		AccountID: accountID,
		Limit: int32(perPage),
		Offset: int32(offset),
	})
	if err != nil {
		return nil,err
	}

	var entriesEntity = []account_entity.Entry{}
	for _,model := range entriesModel {
		entity := account_entity.Entry{
			ID: model.ID,
			AccountID: model.AccountID,
			Amount: int(model.Amount),
			TransactionType: model.TransactionType,
			CreatedAt: model.CreatedAt,
		}
		entriesEntity = append(entriesEntity, entity)
	}

	return entriesEntity,nil
}

func (r *EntryRepository) Get(ctx context.Context, id string) (*account_entity.Entry, error) {
	entryModel,err := r.Queries.GetEntry(ctx,id)
	if err != nil {
		return nil,errors.New("entry not found")
	}

	entryEntity := &account_entity.Entry{
		ID: entryModel.ID,
		AccountID: entryModel.AccountID,
		Amount: int(entryModel.Amount),
		TransactionType: entryModel.TransactionType,
		CreatedAt: entryModel.CreatedAt,
	}

	return entryEntity,nil
}

func (r *EntryRepository) Delete(ctx context.Context, id string) error {
	err := r.Queries.DeleteEntry(ctx,id)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
