package schemas

import (
	"fmt"
)

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Status   string `json:"-"`
}

func (u *UserInput) ValidateUserInput() error {
	if u.Name == "" {
		return fmt.Errorf("the following fields are required: %s", "Name")
	}
	if u.Password == "" {
		return fmt.Errorf("the following fields are required: %s", "Password")
	}
	return nil
}

type User struct {
	Id           int    `db:"id"`
	Name         string `db:"name"`
	PasswordHash string `db:"password_hash"`
	Status       string `db:"status"`
}
