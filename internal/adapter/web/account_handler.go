package web

import (
	dto_account "bank_server/internal/account/application/dto"
	"encoding/json"
	"net/http"
)

// CreateAccountHandler godoc
// @Summary      Create account
// @Description  create account
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        request body   dto_account.InputCreateAccountDto  true  "account request"
// @Success      201  {object}  dto_account.OutputCreateAccountDto
// @Failure      400  {object}  JsonResponse
// @Failure      500  {object}  JsonResponse
// @Router       /account [post]
func (app *Application) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputCreateAccountDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	out, err := app.CreateAccountUseCase.Execute(r.Context(), inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	app.writeJson(w, http.StatusCreated, out)
}

// CreditHandler godoc
// @Summary      Credit value
// @Description  Credit value on account
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        request body   dto_account.InputCreditValueUseCase  true  "account request"
// @Success      201  {object}  dto_account.OutputCreditValueUseCase
// @Failure      400  {object}  JsonResponse
// @Failure      500  {object}  JsonResponse
// @Router       /account/credit [post]
func (app *Application) CreditValueHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputCreditValueUseCase
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	out, err := app.CreditValueUseCase.Execute(r.Context(), inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}
	app.writeJson(w, http.StatusOK, out)
}

// DebitHandler godoc
// @Summary      Credit value
// @Description  Credit value on account
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        request body   dto_account.InputDebitValueUseCase  true  "account request"
// @Success      201  {object}  dto_account.OutputDebitValueUseCase
// @Failure      400  {object}  JsonResponse
// @Failure      500  {object}  JsonResponse
// @Router       /account/debit [post]
func (app *Application) DebitValueHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputDebitValueUseCase
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	out, err := app.DebitValueUseCase.Execute(r.Context(), inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}
	app.writeJson(w, http.StatusOK, out)
}

// DebitHandler godoc
// @Summary      Credit value
// @Description  Credit value on account
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        request body   dto_account.InputTransferUseCase  true  "account request"
// @Success      201  {object}  dto_account.OutputTransferUseCase
// @Failure      400  {object}  JsonResponse
// @Failure      500  {object}  JsonResponse
// @Router       /account/transfer [post]
func (app *Application) TransferHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto_account.InputTransferUseCase
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.TransferHandlerErrosObserver.Inc()
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	
	out, err := app.TransferUseCase.Execute(r.Context(), inputDto)
	if err != nil {
		app.TransferHandlerErrosObserver.Inc()
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}
	app.writeJson(w, http.StatusOK, out)
}
