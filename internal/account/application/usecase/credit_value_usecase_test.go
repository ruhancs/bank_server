package account_usecase_test

import (
	dto_account "bank_server/internal/account/application/dto"
	account_usecase "bank_server/internal/account/application/usecase"
	mock_gateway_account "bank_server/internal/account/application/usecase/mock"
	"bank_server/internal/account/domain/entity"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreditValueUseCase(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	unitOfWorkMock := mock_gateway_account.NewMockUowInterface(ctrl)
	//accountMockRepo := mock_gateway_account.NewMockAccountRepositoryInterface(ctrl)
	//entryMockRepo := mock_gateway_account.NewMockEntryrepositoryInterface(ctrl)

	account,_ := account_entity.NewAccount("user1")
	//entry,_ := entity.NewEntry(account.ID,account_entity.CREDIT,20)

	//unitOfWorkMock.EXPECT().GetRepository(gomock.Any(),gomock.Any()).Return(accountMockRepo,nil)
	//accountMockRepo.EXPECT().GetToUpdate(gomock.Any()).Return(account,nil)
	//unitOfWorkMock.EXPECT().GetRepository(gomock.Any(),gomock.Any()).Return(entryMockRepo,nil)
	//entryMockRepo.EXPECT().Create(entry).Return(nil)
	unitOfWorkMock.EXPECT().Do(gomock.Any(),gomock.Any()).Return(nil)

	creditValueUseCase := account_usecase.NewCreditValueUseCase(unitOfWorkMock)
	input := dto_account.InputCreditValueUseCase{
		AccountID: account.ID,
		Value: 20,
	}
	out,err := creditValueUseCase.Execute(context.Background(),input)

	assert.Nil(t,err)
	assert.NotNil(t,out)
}