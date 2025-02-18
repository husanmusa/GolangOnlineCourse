package handler

import (
	"auth/model"
	pb "auth/proto"
	"context"
)

func (h *AuthController) SignUp(ctx context.Context, req *pb.SignInReq) (*pb.Empty, error) {
	//var newUser model.CreateUserDTO

	newUser := model.CreateUserDTO{Email: req.Email, Password: req.Password}

	if err := h.services.SignUp(newUser); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *AuthController) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	credentials := model.SignInDTO{Email: req.Email, Password: req.Password}
	accessToken, err := h.services.SignIn(credentials)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResp{
		Token: accessToken,
	}, nil
}
