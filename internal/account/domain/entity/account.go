package account_entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Owner     string
	Balance   int
	CreatedAt time.Time
}

func NewAccount(owner string) (*Account,error) {
	account := &Account{
		ID: uuid.NewString(),
		Owner: owner,
		Balance: 0,
		CreatedAt: time.Now(),
	}
	err := account.validate()
	if err != nil {
		return nil,err
	}
	return account,nil
}

func(a *Account) CreditBalance(value int) {
	a.Balance = a.Balance + value
}

func(a *Account) DebitBalance(value int) error {
	if a.Balance - value < 0 {
		return errors.New("balance insufficient")
	} 
	a.Balance = a.Balance - value
	return nil
}

func(a *Account) validate() error {
	if a.Owner == "" {
		return errors.New("owner should not be empty")
	}
	if a.Balance > 0 || a.Balance < 0 {
		return errors.New("invalid balance to init an account")
	}

	return nil
}
