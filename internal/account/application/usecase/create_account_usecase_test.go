package account_usecase_test

import (
	dto_account "bank_server/internal/account/application/dto"
	account_usecase "bank_server/internal/account/application/usecase"
	mock_gateway_account "bank_server/internal/account/application/usecase/mock"
	mock_gateway_user "bank_server/internal/user/application/usecase/mock"
	"bank_server/internal/user/domain/entity"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountUseCase(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	accountMockRepo := mock_gateway_account.NewMockAccountRepositoryInterface(ctrl)
	userMockRepo := mock_gateway_user.NewMockUserRepositoryInterface(ctrl)

	user,_ := entity.NewUser("user1","user@email.com")
	userMockRepo.EXPECT().GetByUserName(gomock.Any(),gomock.Any()).Return(user,nil)
	accountMockRepo.EXPECT().Create(gomock.Any(),gomock.Any()).Return(nil)

	accountUsecase := account_usecase.NewCreateAccountUseCase(accountMockRepo,userMockRepo)
	input := dto_account.InputCreateAccountDto{
		Owner: "user1",
	}
	output,err := accountUsecase.Execute(context.Background(),input)

	assert.Nil(t,err)
	assert.NotNil(t,output)
}