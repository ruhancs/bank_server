package web_user_test

/*
import (
	user_dto "bank_server/internal/user/application/dto"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var userRepo = initUserRepository()

func TestCreateUserHandler(t *testing.T) {
	createUserUseCase := initCreateUserUseCase(userRepo)
	getUserUseCase := initGetUserUseCase(userRepo)
	app := initApplication(createUserUseCase,getUserUseCase)

	go app.Server()

	client := http.Client{}
	input := bytes.NewBuffer([]byte(`{"username": "user1", "email": "user1@email.com"}`))
	resp, err := client.Post("http://localhost:8000/user", "application/json", input)
	assert.Nil(t, err)
	defer resp.Body.Close()

	outputBytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.NotNil(t, string(outputBytes))
	var output user_dto.InputCreateUserDto
	err = json.Unmarshal(outputBytes, &output)
	assert.Nil(t, err)

	assert.NotNil(t,output)
	assert.Equal(t,"user1",output.UserName)
	assert.Equal(t,"user1@email.com",output.Email)
	
	userRepo.Delete(context.Background(),output.UserName)
}

func TestGetUserHandler(t *testing.T) {
	createUserUseCase := initCreateUserUseCase(userRepo)
	getUserUseCase := initGetUserUseCase(userRepo)
	app := initApplication(createUserUseCase,getUserUseCase)

	userEntity,_ := entity.NewUser("user1","user1@email.com")
	userRepo.Create(context.Background(),userEntity)

	go app.Server()

	url := fmt.Sprintf("http://localhost:8000/user/%s","user1")
	client := http.Client{}
	resp,err := client.Get(url)
	assert.Nil(t,err)
	defer resp.Body.Close()

	outputBytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.NotNil(t, string(outputBytes))
	var output user_dto.OutputGetUserDto
	err = json.Unmarshal(outputBytes, &output)
	assert.Nil(t, err)

	assert.NotNil(t,output)
	assert.Equal(t,"user1",output.UserName)
	assert.Equal(t,"user1@email.com",output.Email)

	userRepo.Delete(context.Background(),"user1")
}
*/