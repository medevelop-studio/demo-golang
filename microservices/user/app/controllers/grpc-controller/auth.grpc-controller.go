package grpc_controller

import (
	"context"

	"user/app/services"
	pb "user/proto"
)

type AuthServiceController struct {
	pb.UnimplementedAuthServiceServer
	services *services.Services
}

func (controller *AuthServiceController) Login(
	ctx context.Context,
	data *pb.LoginRequest,
) (*pb.LoginResponse, error) {
	loginData, err := controller.services.AuthService.LoginAdmin(data.Login, data.Password)
	if err != nil {
		return nil, err
	}

	res := &pb.LoginResponse{
		User: &pb.User{
			Id:    loginData.Id,
			Login: loginData.Login,
			Role:  pb.UserRoles(loginData.Role),
		},
		Token: loginData.Token,
	}

	return res, nil
}
