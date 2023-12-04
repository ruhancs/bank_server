package dto_account

import "time"

type InputCreateAccountDto struct {
	Owner string `json:"owner"`
}

type OutputCreateAccountDto struct {
	ID       string    `json:"id"`
	Owner    string    `json:"owner"`
	Balance  int       `json:"balance"`
	CreateAt time.Time `json:"created_at"`
}

type InputCreditValueUseCase struct {
	AccountID string `json:"account_id"`
	Value        int    `json:"value"`
}

type OutputCreditValueUseCase struct {
	Status string `json:"status"`
	Value  int    `json:"value"`
}

type InputDebitValueUseCase struct {
	AccountID string `json:"account_id"`
	Value     int    `json:"value"`
}

type OutputDebitValueUseCase struct {
	Status string `json:"status"`
	Value  int    `json:"value"`
}

type InputTransferUseCase struct {
	FromAccountID string `json:"from_account"`
	ToAccountID   string `json:"to_account"`
	Value         int    `json:"value"`
}

type OutputTransferUseCase struct {
	FromAccountID string `json:"from_account"`
	ToAccountID   string `json:"to_account"`
	Status        string `json:"status"`
	Value         int    `json:"value"`
}
