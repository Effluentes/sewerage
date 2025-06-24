package controller

import (
	"sewerage/internal/domain/dto"
	"sewerage/internal/domain/models"
	"sewerage/internal/domain/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) CreateUser(req dto.CreateUserRequest) (models.User, error) {
	user := models.User{
		Email: req.Email,
	}
	return c.service.CreateUser(&user)
}