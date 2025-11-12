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
	existingUser,err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil && err != sql.ErrNoRows{
		return nil,err // loi database
	}
	if existingUser{
		return nil, fmt.Errorf("username da ton tai")
	}

	existingEmail,err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows{
		return nil,err // loi database
	}
	if existingEmail{
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
