package web

import (
	user_dto "bank_server/internal/user/application/dto"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func(app *Application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto user_dto.InputCreateUserDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,errors.New("invalid input data"), http.StatusBadRequest)
		return
	}
	
	out,err := app.CreateUserUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err, http.StatusBadRequest)
		return
	}

	app.writeJson(w,http.StatusCreated,out)
}

func(app *Application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r,"username")
	fmt.Println("USERNAME")
	fmt.Println(username)

	out,err := app.GetUserUseCase.Execute(r.Context(),username)
	if err != nil {
		fmt.Println(err)
		app.errorJson(w,err,http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,out)
}