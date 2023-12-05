package web

import (
	"bank_server/internal/user/application/dto"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateUserHandler godoc
// @Summary      Create user
// @Description  create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body   user_dto.InputCreateUserDto  true  "user request"
// @Success      201  {object}  user_dto.OutputCreateUserDto
// @Failure      400  {object}  JsonResponse
// @Failure      500  {object}  JsonResponse
// @Router       /user [post]
func (app *Application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto user_dto.InputCreateUserDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w, errors.New("invalid input data"), http.StatusBadRequest)
		return
	}

	out, err := app.CreateUserUseCase.Execute(r.Context(), inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	app.writeJson(w, http.StatusCreated, out)
}

// GetUserHandler godoc
// @Summary      Get user
// @Description  Get user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "user id" Format(uuid)
// @Success      200  {object}  user_dto.OutputGetUserDto
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /user/{id} [get]
func (app *Application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	fmt.Println("USERNAME")
	fmt.Println(username)

	out, err := app.GetUserUseCase.Execute(r.Context(), username)
	if err != nil {
		fmt.Println(err)
		app.errorJson(w, err, http.StatusNotFound)
		return
	}

	app.writeJson(w, http.StatusOK, out)
}
