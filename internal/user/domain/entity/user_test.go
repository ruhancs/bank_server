package entity_test

import (
	"bank_server/internal/user/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("user1", "user@email.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "user1", user.Username)
	assert.Equal(t, "user@email.com", user.Email)
}

func TestNewUserWithInvalidEmail(t *testing.T) {
	user, err := entity.NewUser("user1", "useremail.com")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid email",err.Error())
	assert.Nil(t, user)
}
