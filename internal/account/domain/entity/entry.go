package account_entity

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

const (
	CREDIT = "credit"
	DEBIT  = "debit"
)

type Entry struct {
	ID              string
	AccountID       string
	Amount          int
	TransactionType string
	CreatedAt       time.Time
}

func NewEntry(accountID, transactionType string, amount int) (*Entry, error) {
	entry := &Entry{
		ID:              xid.New().String(),
		AccountID:       accountID,
		Amount:          amount,
		TransactionType: transactionType,
		CreatedAt:       time.Now(),
	}

	err := entry.validate()
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (e *Entry) validate() error {
	if e.AccountID == "" {
		return errors.New("inavalid account id")
	}
	if e.Amount <= 0 {
		return errors.New("inavalid amount")
	}
	return nil
}
