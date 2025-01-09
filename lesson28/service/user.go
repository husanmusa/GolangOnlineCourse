package service

import (
	"lesson28/repository"

	"lesson28/models"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (u *UserService) CreateUser(user models.User) error {
	return u.UserRepo.CreateUser(user)
}
