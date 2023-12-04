package account_repository_test

import (
	account_entity "bank_server/internal/account/domain/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTransfer(fromAccount, toAccount string) *account_entity.Transfer {
	transfer, _ := account_entity.NewTransfer(fromAccount, toAccount, 10)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	ctx := context.Background()
	user1 := createUser("user1","user1@email.com")
	user2 := createUser("user2","user2@email.com")
	fromAccount := createAccount(user1.Username)
	ToAccount := createAccount(user2.Username)
	transfer := createTransfer(fromAccount.ID, ToAccount.ID)
	userRepository, accountRepository, _, transferRepository := initRepositories()
	userRepository.Create(ctx, user1)
	userRepository.Create(ctx, user2)
	accountRepository.Create(ctx, fromAccount)
	accountRepository.Create(ctx, ToAccount)

	err := transferRepository.Create(ctx, transfer)

	assert.Nil(t, err)

	transferRepository.Delete(ctx, transfer.ID)
	accountRepository.Delete(ctx, fromAccount.ID)
	accountRepository.Delete(ctx, ToAccount.ID)
	userRepository.Delete(ctx, user1.Username)
	userRepository.Delete(ctx, user2.Username)
}

func TestListTransfer(t *testing.T) {
	ctx := context.Background()
	user1 := createUser("user1","user1@email.com")
	user2 := createUser("user2","user2@email.com")
	fromAccount := createAccount(user1.Username)
	ToAccount := createAccount(user2.Username)
	transfer1 := createTransfer(fromAccount.ID, ToAccount.ID)
	transfer2 := createTransfer(ToAccount.ID, fromAccount.ID)
	userRepository, accountRepository, _, transferRepository := initRepositories()
	userRepository.Create(ctx, user1)
	userRepository.Create(ctx, user2)
	accountRepository.Create(ctx, fromAccount)
	accountRepository.Create(ctx, ToAccount)

	transferRepository.Create(ctx, transfer1)
	transferRepository.Create(ctx, transfer2)

	transfers,err := transferRepository.List(ctx,1,1)

	assert.Nil(t,err)
	assert.NotNil(t,transfers)
	assert.Equal(t,1,len(transfers))
	assert.Equal(t,transfer1.ID,transfers[0].ID)
	
	transfers,err = transferRepository.List(ctx,2,1)

	assert.Nil(t,err)
	assert.NotNil(t,transfers)
	assert.Equal(t,2,len(transfers))
	assert.Equal(t,transfer1.ID,transfers[0].ID)
	assert.Equal(t,transfer2.ID,transfers[1].ID)

	transferRepository.Delete(ctx, transfer1.ID)
	transferRepository.Delete(ctx, transfer2.ID)
	accountRepository.Delete(ctx, fromAccount.ID)
	accountRepository.Delete(ctx, ToAccount.ID)
	userRepository.Delete(ctx, user1.Username)
	userRepository.Delete(ctx, user2.Username)
}

func TestGetTransfer(t *testing.T) {
	ctx := context.Background()
	user1 := createUser("user1","user1@email.com")
	user2 := createUser("user2","user2@email.com")
	fromAccount := createAccount(user1.Username)
	ToAccount := createAccount(user2.Username)
	transfer := createTransfer(fromAccount.ID, ToAccount.ID)
	userRepository, accountRepository, _, transferRepository := initRepositories()
	userRepository.Create(ctx, user1)
	userRepository.Create(ctx, user2)
	accountRepository.Create(ctx, fromAccount)
	accountRepository.Create(ctx, ToAccount)

	transferRepository.Create(ctx, transfer)

	transferFounded,err := transferRepository.Get(ctx,transfer.ID)

	assert.Nil(t, err)
	assert.NotNil(t,transfer)
	assert.Equal(t,transfer.ID,transferFounded.ID)
	assert.Equal(t,transfer.FromAccountID,transferFounded.FromAccountID)
	assert.Equal(t,transfer.ToAccountID,transferFounded.ToAccountID)
	assert.Equal(t,transfer.Amount,transferFounded.Amount)

	transferRepository.Delete(ctx, transfer.ID)
	accountRepository.Delete(ctx, fromAccount.ID)
	accountRepository.Delete(ctx, ToAccount.ID)
	userRepository.Delete(ctx, user1.Username)
	userRepository.Delete(ctx, user2.Username)
}
