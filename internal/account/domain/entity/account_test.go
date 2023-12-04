package account_entity_test

import (
	account_entity "bank_server/internal/account/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	account,err := account_entity.NewAccount("user1")

	assert.Nil(t,err)
	assert.NotNil(t,account)
	assert.Equal(t,"user1",account.Owner)
	assert.Equal(t,0,account.Balance)
}

func TestNewAccountWithInvalidOwner(t *testing.T) {
	account,err := account_entity.NewAccount("")

	assert.Nil(t,account)
	assert.NotNil(t,err)
	assert.Equal(t,"owner should not be empty",err.Error())
}

func TestCreditBalance(t *testing.T) {
	account,_ := account_entity.NewAccount("user1")

	account.CreditBalance(20)

	assert.Equal(t,20,account.Balance)
}

func TestDebitBalance(t *testing.T) {
	account,_ := account_entity.NewAccount("user1")

	err := account.DebitBalance(20)
	
	assert.NotNil(t,"balance insufficient",err.Error())
	
	account.CreditBalance(30)
	err = account.DebitBalance(30)

	assert.Nil(t,err)
	assert.Equal(t,0,account.Balance)
	
	account.CreditBalance(30)
	err = account.DebitBalance(20)
	
	assert.Nil(t,err)
	assert.Equal(t,10,account.Balance)
}
