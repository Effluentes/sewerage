package dto

import (
	"fmt"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required`
}

func (request *CreateUserRequest) Validate() error {
	if request.Email == "" {
		return fmt.Errorf("email is required")
	}
	if request.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}