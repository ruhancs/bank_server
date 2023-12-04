package account_repository_test

import (
	account_entity "bank_server/internal/account/domain/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createEntry(accountID,trType string) *account_entity.Entry {
	entry,_ := account_entity.NewEntry(accountID,trType,20)
	return entry
}

func TestCreateEntry(t *testing.T) {
	ctx := context.Background()
	user := createUser("user1","user@email.com")
	account := createAccount(user.Username)
	entry := createEntry(account.ID,account_entity.CREDIT)
	userRepository,accountRepository,entryRepository,_ := initRepositories()
	userRepository.Create(ctx,user)
	accountRepository.Create(ctx,account)

	err := entryRepository.Create(ctx,entry)

	assert.Nil(t,err)
	
	entryRepository.Delete(ctx,entry.ID)
	accountRepository.Delete(ctx,account.ID)
	userRepository.Delete(ctx,user.Username)
}

func TestListEntries(t *testing.T) {
	ctx := context.Background()
	user := createUser("user1","user@email.com")
	account := createAccount(user.Username)
	entry1 := createEntry(account.ID,account_entity.CREDIT)
	entry2 := createEntry(account.ID,account_entity.CREDIT)
	userRepository,accountRepository,entryRepository,_ := initRepositories()
	userRepository.Create(ctx,user)
	accountRepository.Create(ctx,account)

	entryRepository.Create(ctx,entry1)
	entryRepository.Create(ctx,entry2)

	entries,err := entryRepository.List(ctx,account.ID,1,1)

	assert.Nil(t,err)
	assert.NotNil(t,entries)
	assert.Equal(t,1,len(entries))
	assert.Equal(t,entry1.AccountID,entries[0].AccountID)

	entries,err = entryRepository.List(ctx,account.ID,2,1)

	assert.Nil(t,err)
	assert.NotNil(t,entries)
	assert.Equal(t,2,len(entries))
	assert.Equal(t,entry1.AccountID,entries[0].AccountID)
	assert.Equal(t,entry1.Amount,entries[0].Amount)
	assert.Equal(t,entry2.AccountID,entries[1].AccountID)
	assert.Equal(t,entry2.Amount,entries[1].Amount)
	
	entryRepository.Delete(ctx,entry1.ID)
	entryRepository.Delete(ctx,entry2.ID)
	accountRepository.Delete(ctx,account.ID)
	userRepository.Delete(ctx,user.Username)
}

func TestGetEntry(t *testing.T) {
	ctx := context.Background()
	user := createUser("user1","user@email.com")
	account := createAccount(user.Username)
	entry := createEntry(account.ID,account_entity.CREDIT)
	userRepository,accountRepository,entryRepository,_ := initRepositories()
	userRepository.Create(ctx,user)
	accountRepository.Create(ctx,account)

	entryRepository.Create(ctx,entry)

	entryFounded,err := entryRepository.Get(ctx,entry.ID)

	assert.Nil(t,err)
	assert.NotNil(t,entryFounded)
	assert.Equal(t,entry.AccountID,entryFounded.AccountID)
	assert.Equal(t,entry.Amount,entryFounded.Amount)
	assert.Equal(t,entry.ID,entryFounded.ID)
	assert.Equal(t,entry.TransactionType,entryFounded.TransactionType)
	
	entryRepository.Delete(ctx,entry.ID)
	accountRepository.Delete(ctx,account.ID)
	userRepository.Delete(ctx,user.Username)
}