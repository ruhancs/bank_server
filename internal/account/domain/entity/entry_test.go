package account_entity_test

import (
	"bank_server/internal/account/domain/entity"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEntry(t *testing.T) {
	accountID := uuid.NewV4().String()
	entrie, err := account_entity.NewEntry(accountID, account_entity.CREDIT, 10)

	assert.Nil(t, err)
	assert.NotNil(t, entrie)
	assert.Equal(t, accountID, entrie.AccountID)
	assert.Equal(t, 10, entrie.Amount)
}

func TestNewEntryWithInvalidAmount(t *testing.T) {
	accountID := uuid.NewV4().String()
	entrie, err := account_entity.NewEntry(accountID, account_entity.DEBIT, -1)

	assert.Nil(t, entrie)
	assert.NotNil(t, err)
	assert.Equal(t, "inavalid amount", err.Error())
}
