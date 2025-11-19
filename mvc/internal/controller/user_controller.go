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

func (c *UserController) GetUserByIDController(id int) (*dto.UserReponse, error) {
	userResponse, err := c.UserSev.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}

func (c *UserController) GetAllUserController() ([]dto.UserReponse, error) {
	usersResponse, err := c.UserSev.GetAllUser()
	if err != nil {
		return nil, err
	}
	return usersResponse, nil
}
func (c *UserController) UpdateUserByIDController(req dto.UpdateUserRequest, id int) (*dto.UserReponse, error) {
	userResponse, err := c.UserSev.UpdateUserByID(req, id)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}
func (c *UserController) DeleteUserByIDController(id int) error {
	err := c.UserSev.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}