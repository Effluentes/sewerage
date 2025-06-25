package controller

import (
	"sewerage/internal/domain/dto"
	"sewerage/internal/domain/models"
	"sewerage/internal/domain/services"
	"net/mail"
	"fmt"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) CreateUser(req *dto.CreateUserRequest) error {
	if err := validateCreateUserRequest(req); err != nil {
		return err
	}
	user := models.User{
		Email: req.Email,
	}
	return c.service.CreateUser(&user)
}

func validEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func validateCreateUserRequest(request *dto.CreateUserRequest) error {
	if request.Email == "" {
		return fmt.Errorf("email is required")
	}
	if request.Password == "" {
		return fmt.Errorf("password is required")
	}
	if !validEmail(request.Email) {
		return fmt.Errorf("email is invalid")
	}
	return nil
}