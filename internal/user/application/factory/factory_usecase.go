package user_factory

import (
	user_usecase "bank_server/internal/user/application/usecase"
	user_repository "bank_server/internal/user/infra/repository"
	"database/sql"

	"go.uber.org/zap"
)


func CreateUserUseCase(db *sql.DB,logger *zap.Logger) *user_usecase.CreateUserUseCase {
	userRepository := user_repository.NewUserRepository(db)
	usecase := user_usecase.NewCreateUserUseCase(userRepository,logger)
	return usecase
}

func GetUserUseCase(db *sql.DB, logger *zap.Logger) *user_usecase.GetUserUseCase {
	userRepository := user_repository.NewUserRepository(db)
	usecase := user_usecase.NewGetUserUseCase(userRepository,logger)
	return usecase
}