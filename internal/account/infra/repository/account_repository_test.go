package account_repository_test

import (
	account_entity "bank_server/internal/account/domain/entity"
	account_repository "bank_server/internal/account/infra/repository"
	"bank_server/internal/user/domain/entity"
	user_repository "bank_server/internal/user/infra/repository"
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	_ "github.com/lib/pq"
)

func createUser(username,email string) *entity.User {
	user, _ := entity.NewUser(username, email)
	return user
}

func createAccount(owner string) *account_entity.Account  {
	account, _ := account_entity.NewAccount(owner)
	return account
}

func initRepositories() (
	*user_repository.UserRepository, 
	*account_repository.AccountRepository, 
	*account_repository.EntryRepository,
	*account_repository.TransferRepository,
) {
	err := godotenv.Load("../../../../.envexample")
	if err != nil {
		panic(err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println(err)
		log.Fatal("cannot connect to db")
	}
	// TODO fechar conexao com DB
	//defer dbConn.Close()

	userRepository := user_repository.NewUserRepository(dbConn)
	accountRepository := account_repository.NewAccountRepository(dbConn)
	entryRepository := account_repository.NewEntryRepository(dbConn)
	transferRepository := account_repository.NewTransferRepository(dbConn)
	return userRepository, accountRepository,entryRepository,transferRepository
}

func TestCreateAccount(t *testing.T) {
	ctx := context.Background()
	user := createUser("user1","user@email.com")
	account := createAccount(user.Username)
	userRepository,accountRepository,_,_ := initRepositories()
	userRepository.Create(ctx,user)

	err := accountRepository.Create(ctx,account)

	assert.Nil(t,err)
	accountRepository.Delete(ctx,account.ID)
}

func TestGetAccount(t *testing.T) {
	ctx := context.Background()
	user := createUser("user1","user@email.com")
	account := createAccount(user.Username)
	userRepository,accountRepository,_,_ := initRepositories()
	userRepository.Create(ctx,user)

	err := accountRepository.Create(ctx,account)
	assert.Nil(t,err)
	
	accountFounded,err := accountRepository.Get(ctx,account.ID)
	
	assert.Nil(t,err)
	assert.NotNil(t,accountFounded)
	assert.Equal(t,account.ID,accountFounded.ID)
	assert.Equal(t,account.Balance,accountFounded.Balance)
	assert.Equal(t,account.Owner,accountFounded.Owner)
	
	accountRepository.Delete(ctx,account.ID)
}

func TestGetAccountToUpdate(t *testing.T) {
	ctx := context.Background()
	user := createUser("user1","user@email.com")
	account := createAccount(user.Username)
	userRepository,accountRepository,_,_ := initRepositories()
	userRepository.Create(ctx,user)

	err := accountRepository.Create(ctx,account)
	assert.Nil(t,err)
	
	accountFounded,err := accountRepository.GetToUpdate(ctx,account.ID)
	
	assert.Nil(t,err)
	assert.NotNil(t,accountFounded)
	assert.Equal(t,account.ID,accountFounded.ID)
	assert.Equal(t,account.Balance,accountFounded.Balance)
	assert.Equal(t,account.Owner,accountFounded.Owner)
	
	accountRepository.Delete(ctx,account.ID)
}

func TestUpdateAccountBalance(t *testing.T) {
	ctx := context.Background()
	user := createUser("user1","user@email.com")
	account := createAccount(user.Username)
	userRepository,accountRepository,_,_ := initRepositories()
	userRepository.Create(ctx,user)

	err := accountRepository.Create(ctx,account)
	assert.Nil(t,err)
	
	err = accountRepository.UpdateBalance(ctx,account.ID,30)
	assert.Nil(t,err)
	
	accountFounded,err := accountRepository.GetToUpdate(ctx,account.ID)
	
	assert.Nil(t,err)
	assert.NotNil(t,accountFounded)
	assert.Equal(t,account.ID,accountFounded.ID)
	assert.Equal(t,30,accountFounded.Balance)
	assert.Equal(t,account.Owner,accountFounded.Owner)
	
	accountRepository.Delete(ctx,account.ID)
}
