package account_repository

import (
	account_entity "bank_server/internal/account/domain/entity"
	"bank_server/sql/db"
	"context"
	"database/sql"
	"errors"
)

type TransferRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewTransferRepository(database *sql.DB) *TransferRepository{
	return &TransferRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(r *TransferRepository) Create(ctx context.Context,transfer *account_entity.Transfer) error {
	err := r.Queries.CreateTransfer(ctx,db.CreateTransferParams{
		ID: transfer.ID,
		FromAccountID: transfer.FromAccountID,
		ToAccountID: transfer.ToAccountID,
		Amount: int64(transfer.Amount),
	})
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func(r *TransferRepository) List(ctx context.Context,perPage,page int) ([]account_entity.Transfer,error){
	offset := (page - 1) * perPage
	transferModels,err := r.Queries.ListTransfers(ctx,db.ListTransfersParams{
		Limit: int32(perPage),
		Offset: int32(offset),
	})
	if err != nil {
		return nil,err
	}

	var transferEntities = []account_entity.Transfer{}
	for _,model := range transferModels{
		entity := account_entity.Transfer{
			ID: model.ID,
			FromAccountID: model.FromAccountID,
			ToAccountID: model.ToAccountID,
			Amount: int(model.Amount),
			CreatedAt: model.CreatedAt,
		}
		transferEntities = append(transferEntities, entity)
	}

	return transferEntities,nil
}

func(r *TransferRepository) Get(ctx context.Context,id string) (*account_entity.Transfer,error){
	tranferModel,err := r.Queries.GetTransfer(ctx,id)
	if err != nil {
		return nil,errors.New("transfer not found")
	}
	
	transfer := account_entity.Transfer{
		ID: tranferModel.ID,
		FromAccountID: tranferModel.FromAccountID,
		ToAccountID: tranferModel.ToAccountID,
		Amount: int(tranferModel.Amount),
		CreatedAt: tranferModel.CreatedAt,
	}

	return &transfer,nil
}

func(r *TransferRepository) Delete(ctx context.Context,id string) error {
	err := r.Queries.DeleteTransfer(ctx,id)
	if err != nil {
		return errors.New("error to delete transaction")
	}

	return nil
}

