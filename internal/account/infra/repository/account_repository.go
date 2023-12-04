package account_repository

import (
	"bank_server/internal/account/domain/entity"
	"bank_server/sql/db"
	"context"
	"database/sql"
	"errors"
)

type AccountRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewAccountRepository(database *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(r *AccountRepository) Create(ctx context.Context, account *account_entity.Account) error {
	err := r.Queries.CreateAccount(ctx,db.CreateAccountParams{
		ID: account.ID,
		Owner: account.Owner,
		Balance: int64(account.Balance),
	})
	if err != nil {
		return errors.New("failed to create account")
	}
	return nil
}

func(r *AccountRepository) Get(ctx context.Context, id string) (*account_entity.Account, error) {
	accountModel,err := r.Queries.GetAccount(ctx,id)
	if err != nil {
		return nil,errors.New("account not found")
	}

	accountEntity := &account_entity.Account{
		ID: accountModel.ID,
		Owner: accountModel.Owner,
		Balance: int(accountModel.Balance),
		CreatedAt: accountModel.CreatedAt,
	}

	return accountEntity,nil
}

func(r *AccountRepository) GetToUpdate(ctx context.Context, id string) (*account_entity.Account, error) {
	accountModel,err := r.Queries.GetAccountForUpdate(ctx,id)
	if err != nil {
		return nil,errors.New("account not found")
	}

	accountEntity := &account_entity.Account{
		ID: accountModel.ID,
		Owner: accountModel.Owner,
		Balance: int(accountModel.Balance),
		CreatedAt: accountModel.CreatedAt,
	}

	return accountEntity,nil
}

func(r *AccountRepository) UpdateBalance(ctx context.Context, id string, balance int) error {
	err := r.Queries.UpdateAccount(ctx,db.UpdateAccountParams{
		ID: id,
		Balance: int64(balance),
	})
	if err != nil {
		return errors.New("error to update account balance")
	}


	return nil
}

func(r *AccountRepository) Delete(ctx context.Context,accountID string) error {
	err := r.Queries.DeleteAccount(ctx,accountID)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}