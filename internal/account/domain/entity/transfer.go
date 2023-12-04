package account_entity

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Transfer struct {
	ID            string
	FromAccountID string
	ToAccountID   string
	Amount        int
	CreatedAt     time.Time
}

func NewTransfer(fromAccountID,toAccountID string, amount int) (*Transfer,error) {
	transfer := &Transfer{
		ID: xid.New().String(),
		FromAccountID: fromAccountID,
		ToAccountID: toAccountID,
		Amount: amount,
		CreatedAt: time.Now(),
	}

	err := transfer.validate()
	if err != nil {
		return nil,err
	}

	return transfer,nil
}

func(t *Transfer) validate() error {
	if t.FromAccountID == "" {
		return errors.New("from account id should not be empty")
	}
	if t.ToAccountID == "" {
		return errors.New("to account id should not be empty")
	}
	if t.Amount <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}
