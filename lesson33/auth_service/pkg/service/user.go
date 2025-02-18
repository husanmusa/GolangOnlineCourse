package service

import (
	"auth/model"
	helper "auth/pkg/helpers"
	"auth/pkg/repository"
	"fmt"
	"log/slog"
	"time"
)

type User interface {
	SignUp(model.CreateUserDTO) error
	SignIn(model.SignInDTO) (string, error)
	GetUserByEmail(string) (model.User, error)
}

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository,
	}
}

func (u *UserService) SignUp(newUser model.CreateUserDTO) error {
	user, _ := u.GetUserByEmail(newUser.Email)
	if user.Email == newUser.Email {
		slog.Error("sign-up: email already exist")
		return fmt.Errorf("user with %s already exist", newUser.Email)
	}
	hashed, err := helper.HashPassword(newUser.Password)
	if err != nil {
		slog.Error("sign-up: could not hash password")
	}
	newUser.Password = hashed
	return u.repository.InsertUser(newUser)
}

func (u *UserService) SignIn(credentials model.SignInDTO) (string, error) {
	foundedUser, err := u.GetUserByEmail(credentials.Email)
	if err != nil {
		slog.Error(fmt.Sprintf("sign-in: %s", err.Error()))
		return "", fmt.Errorf("could not retrieve user by email")
	}

	if credentials.Email != foundedUser.Email {
		slog.Error(fmt.Sprintf("sign-in: %s not found", credentials.Email))
		return "", fmt.Errorf("email %s does not exist", credentials.Email)
	}

	if err := helper.VerifyPassword(credentials.Password, foundedUser.Password); err != nil {
		slog.Error(fmt.Sprintf("sign-in: invalid password: %s", err.Error()))
		return "", fmt.Errorf("invalid password")
	}

	token, err := helper.GenerateAccessToken(foundedUser, 2*time.Hour)
	if err != nil {
		slog.Error(fmt.Sprintf("sign-in: %s", err.Error()))
		return "", fmt.Errorf("could not generate access token")
	}

	return token, nil
}

func (u *UserService) GetUserByEmail(email string) (model.User, error) {
	return u.repository.FindUserByEmail(email)
}
