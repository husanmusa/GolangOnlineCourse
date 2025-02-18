package handler

import (
	_ "auth/docs"
	"auth/pkg/pubsub"
	"auth/pkg/service"
	pb "auth/proto"
)

type AuthController struct {
	services *service.Service
	ps       *pubsub.PubSub
	pb.UnimplementedAuthServiceServer
}

func NewHandler(services *service.Service, ps *pubsub.PubSub) *AuthController {
	return &AuthController{
		services: services,
		ps:       ps,
	}
}
