package user_repository_test

import (
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

func createUser() *entity.User{
	user,_:= entity.NewUser("user1","user@email.com")
	return user 
}


func initUserRepository() *user_repository.UserRepository {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		panic(err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	dbConn,err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Println(err)
		log.Fatal("cannot connect to db")
	}
	// TODO fechar conexao com DB
	//defer dbConn.Close()

	userRepository := user_repository.NewUserRepository(dbConn)
	return userRepository
}

func TestCreateUser(t *testing.T) {
	user := createUser()
	userRepository := initUserRepository()

	err := userRepository.Create(context.Background(),user)

	assert.Nil(t,err)
	userRepository.Delete(context.Background(),user.Username)
}

func TestGetUser(t *testing.T) {
	user := createUser()
	userRepository := initUserRepository()
	
	err := userRepository.Create(context.Background(),user)
	assert.Nil(t,err)

	userFounded,err := userRepository.GetByUserName(context.Background(),user.Username)

	assert.Nil(t,err)
	assert.NotNil(t,userFounded)
	assert.Equal(t,user.Username,userFounded.Username)
	assert.Equal(t,user.Email,userFounded.Email)

	userRepository.Delete(context.Background(),user.Username)
}