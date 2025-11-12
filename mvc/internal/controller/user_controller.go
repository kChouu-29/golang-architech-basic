package controller

import (
	"mvc/internal/dto"
	"mvc/internal/service"
)

type UserController struct {
	UserSev service.UserService
}

func NewUserController(r service.UserService) *UserController {
	return &UserController{
		UserSev: r,
	}
}

func (c *UserController) CreateUserController(req dto.CreateUserRequest) (*dto.UserReponse, error) {
	userReponse, err := c.UserSev.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return userReponse, nil
}

