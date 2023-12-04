package user_repository

import (
	"bank_server/internal/user/domain/entity"
	"bank_server/sql/db"
	"context"
	"database/sql"
	"errors"
)

type UserRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewUserRepository(database *sql.DB) *UserRepository{
	return &UserRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(r *UserRepository) Create(ctx context.Context,user *entity.User) error {
	err := r.Queries.CreateUser(ctx,db.CreateUserParams{
		Username: user.Username,
		Email: user.Email,
	})
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func(r *UserRepository) GetByUserName(ctx context.Context, username string) (*entity.User,error) {
	userModel,err := r.Queries.GetUser(ctx,username)
	if err != nil {
		return nil,errors.New("user not found")
	}
	userEntity := &entity.User{
		Username: userModel.Username,
		Email: userModel.Email,
		CreatedAt: userModel.CreatedAt,
	}

	return userEntity,nil
}

func(r *UserRepository) Delete(ctx context.Context,username string) error {
	err := r.Queries.DeleteUser(ctx,username)
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}