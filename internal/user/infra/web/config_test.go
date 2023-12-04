package web_user_test

import (
	user_usecase "bank_server/internal/user/application/usecase"
	"bank_server/internal/user/domain/entity"
	gateway_user "bank_server/internal/user/domain/gateway"
	user_repository "bank_server/internal/user/infra/repository"
	web_user "bank_server/internal/user/infra/web"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func newUser(username,email string) *entity.User {
	user,_ := entity.NewUser(username,email)
	return user
}

func initUserRepository() *user_repository.UserRepository {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("Error loading .env")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")

	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println(err)
		log.Fatal("cannot connect to db")
	}

	repository := user_repository.NewUserRepository(db)
	return repository
}

func initCreateUserUseCase(repo gateway_user.UserRepositoryInterface) *user_usecase.CreateUserUseCase{
	usecase := user_usecase.NewCreateUserUseCase(repo)

	return usecase
}

func initGetUserUseCase(repo gateway_user.UserRepositoryInterface) *user_usecase.GetUserUseCase{
	usecase := user_usecase.NewGetUserUseCase(repo)

	return usecase
}

func initApplication(createUserUseCase *user_usecase.CreateUserUseCase, getUserUseCase *user_usecase.GetUserUseCase) *web_user.Application{
	app := web_user.NewApplication(createUserUseCase,getUserUseCase)

	return app
}

func clearUserDB(repo *user_repository.UserRepository,username string) {
	repo.Queries.DeleteUser(context.Background(),username)
}