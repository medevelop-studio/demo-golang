package grpc_controller

import (
	"context"

	user_domain "user/app/domain/user"
	"user/app/services"
	"user/proto"
)

type UserServiceController struct {
	proto.UnimplementedUserServiceServer
	services *services.Services
}

func (controller *UserServiceController) GetUserById(
	ctx context.Context,
	data *proto.GetUserByIdRequest,
) (*proto.User, error) {
	user, err := controller.services.UserService.GetUserById(data.Id)
	if err != nil {
		return nil, err
	}

	res := &proto.User{
		Id:    user.Id,
		Login: user.Login,
		Role:  proto.UserRoles(user.Role),
	}

	return res, nil
}

func (controller *UserServiceController) CreateUser(
	ctx context.Context,
	data *proto.CreateUserRequest,
) (*proto.User, error) {
	loginData, err := controller.services.UserService.Create(&user_domain.CreateUserDto{
		Login:    data.Login,
		Password: data.Password,
		Role:     user_domain.UserRole(data.Role),
	})
	if err != nil {
		return nil, err
	}

	res := &proto.User{
		Id:    loginData.Id,
		Login: loginData.Login,
		Role:  proto.UserRoles(loginData.Role),
	}

	return res, nil
}
