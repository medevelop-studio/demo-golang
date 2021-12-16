package user_service

import (
	"context"
	"http/proto"

	"google.golang.org/grpc"
)

type UserService struct {
	userServiceClient proto.UserServiceClient
	ctx               context.Context
}

func Create(userMicroserviceConn *grpc.ClientConn) (*UserService, error) {
	userServiceClient := proto.NewUserServiceClient(userMicroserviceConn)

	userService := &UserService{
		userServiceClient: userServiceClient,
		ctx:               context.Background(),
	}

	return userService, nil
}
