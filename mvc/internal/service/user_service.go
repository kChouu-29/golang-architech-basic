package service

import (
	"database/sql"
	"fmt"
	"mvc/internal/dto"
	"mvc/internal/model"
	"mvc/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (*dto.UserReponse, error)
	GetUserByID(id int) (*dto.UserReponse, error)
	GetAllUser() ([]dto.UserReponse, error)
	UpdateUserByID(req dto.UpdateUserRequest, id int) (*dto.UserReponse, error)
	DeleteUserByID(id int) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (*dto.UserReponse, error) {
	existingUser, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err // loi database
	}
	if existingUser {
		return nil, fmt.Errorf("username da ton tai")
	}

	existingEmail, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err // loi database
	}
	if existingEmail {
		return nil, fmt.Errorf("email da ton tai")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHard), bcrypt.DefaultCost)
	if err != nil {
		return nil, err

	}
	userToCreated := &model.Users{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHard: string(hashedPassword),
		Age:          req.Age,
	}
	createdUser, err := s.userRepo.CreateUser(userToCreated)
	if err != nil {
		return nil, err
	}
	reponse := &dto.UserReponse{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		Email:     createdUser.Email,
		Age:       createdUser.Age,
		CreatedAt: createdUser.CreatedAt.String(),
	}
	return reponse, nil
}

func (s *userService) GetUserByID(id int) (*dto.UserReponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	reponse := &dto.GetUserReponse{
		Message: "Get user successfully",
		Data: &dto.UserReponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Age:       user.Age,
			CreatedAt: user.CreatedAt.String(),
			UpdateAt:  user.UpdatedAt.String(),
		},
	}
	return reponse.Data, nil
}

func (s *userService) GetAllUser() ([]dto.UserReponse, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	var userRes []dto.UserReponse
	for _, user := range users {
		userRes = append(userRes, dto.UserReponse{

			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Age:       user.Age,
			CreatedAt: user.CreatedAt.String(),
			UpdateAt:  user.UpdatedAt.String(),
		})
	}
	return userRes, nil
}
func (s *userService) UpdateUserByID(req dto.UpdateUserRequest, id int) (*dto.UserReponse, error) {
	userToUpdate := &model.Users{
		ID:           id,
		Username:     req.Username,
		Email:        req.Email,
		Age:          req.Age,
	}
	updatedUser, err := s.userRepo.UpdateUserByID(userToUpdate)
	if err != nil {
		return nil, err
	}
	reponse := &dto.UserReponse{
		ID:        updatedUser.ID,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		Age:       updatedUser.Age,
		CreatedAt: updatedUser.CreatedAt.String(),
		UpdateAt:  updatedUser.UpdatedAt.String(),
	}
	return reponse, nil
}
func (s *userService) DeleteUserByID(id int) error {

	err := s.userRepo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}