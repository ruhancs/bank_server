package user_usecase_test

import (
	"bank_server/internal/user/application/dto"
	user_usecase "bank_server/internal/user/application/usecase"
	mock_gateway "bank_server/internal/user/application/usecase/mock"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	userMockRepo := mock_gateway.NewMockUserRepositoryInterface(ctrl)
	input := user_dto.InputCreateUserDto{
		UserName: "user1",
		Email:    "user@email.com",
	}
	userMockRepo.EXPECT().Create(gomock.Any(),gomock.Any()).Return(nil)

	createUserUsecase := user_usecase.NewCreateUserUseCase(userMockRepo)
	out, err := createUserUsecase.Execute(context.Background(),input)

	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.Equal(t, input.UserName, out.UserName)
	assert.Equal(t, input.Email, out.Email)
}

func TestCreateUserUseCaseWithInvalidEmail(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	userMockRepo := mock_gateway.NewMockUserRepositoryInterface(ctrl)
	input := user_dto.InputCreateUserDto{
		UserName: "user1",
		Email:    "user",
	}

	createUserUsecase := user_usecase.NewCreateUserUseCase(userMockRepo)
	out, err := createUserUsecase.Execute(context.Background(),input)

	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.Equal(t,"invalid email",err.Error())
}
