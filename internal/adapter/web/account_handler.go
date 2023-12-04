package web

import (
	dto_account "bank_server/internal/account/application/dto"
	"encoding/json"
	"net/http"
)

func(app *Application) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputCreateAccountDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	out,err := app.CreateAccountUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	app.writeJson(w,http.StatusCreated,out)
}

func(app *Application) CreditValueHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputCreditValueUseCase
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	out,err := app.CreditValueUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusInternalServerError)
		return
	}
	app.writeJson(w,http.StatusOK,out)
}

func(app *Application) DebitValueHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputDebitValueUseCase
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	out,err := app.DebitValueUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusInternalServerError)
		return
	}
	app.writeJson(w,http.StatusOK,out)
}

func(app *Application) TransferHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputTransferUseCase
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	out,err := app.TransferUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusInternalServerError)
		return
	}
	app.writeJson(w,http.StatusOK,out)
}