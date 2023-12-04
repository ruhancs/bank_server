package account_entity_test

import (
	"bank_server/internal/account/domain/entity"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransfer(t *testing.T) {
	fromAccount := xid.New().String()
	toAccount := xid.New().String()
	transfer,err := account_entity.NewTransfer(fromAccount,toAccount,50)

	assert.Nil(t,err)
	assert.NotNil(t,transfer)
	assert.Equal(t,fromAccount,transfer.FromAccountID)
	assert.Equal(t,toAccount,transfer.ToAccountID)
	assert.Equal(t,50,transfer.Amount)
}