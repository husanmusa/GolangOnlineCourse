package service

import (
	"auth/pkg/pubsub"
	"auth/pkg/repository"
)

type Service struct {
	User
}

func NewService(repository *repository.Repository, ps *pubsub.PubSub) *Service {
	return &Service{
		User: NewUserService(repository.User),
	}
}
