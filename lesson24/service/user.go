package service

import (
	"expensity/repository"

	"expensity/models"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (u *UserService) CreateUser(user models.User) error {
	return u.UserRepo.CreateUser()
}
