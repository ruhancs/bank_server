package entity

import (
	"errors"
	"net/mail"
	"time"

)

type User struct {
	//ID        string
	Username  string
	Email     string
	CreatedAt time.Time
}

func NewUser(username, email string) (*User, error) {
	user := &User{
		//ID:        uuid.NewV4().String(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}

	err := user.validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) validate() error {
	if u.Username == "" {
		return errors.New("username should not be empty")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil
}
