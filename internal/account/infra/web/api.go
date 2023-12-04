package account_web

import account_usecase "bank_server/internal/account/application/usecase"

type Application struct {
	CreateAccountUseCase *account_usecase.CreateAccountUseCase
	CreditValueUseCase *account_usecase.CreditValueUseCase
	DebitValueUseCase *account_usecase.DebitValueUseCase
	TransferUseCase *account_usecase.TransferUseCase
}

func NewApplication(
	createAccountUseCase *account_usecase.CreateAccountUseCase,
	creditValuetUseCase *account_usecase.CreditValueUseCase,
	debitValueUseCase *account_usecase.DebitValueUseCase,
	transferUseCase *account_usecase.TransferUseCase,
) *Application {
	return &Application{
		CreateAccountUseCase: createAccountUseCase,
		CreditValueUseCase: creditValuetUseCase,
		DebitValueUseCase: debitValueUseCase,
		TransferUseCase: transferUseCase,
	}
}
