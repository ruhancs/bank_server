package user_dto

import "time"

type InputCreateUserDto struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type OutputCreateUserDto struct {
	//ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type OutputGetUserDto struct {
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
