package user_service

import (
	"chat/proto"
	"context"

	"google.golang.org/grpc"
)

type UserService struct {
	userServiceClient proto.UserServiceClient
	ctx               context.Context
}

type Config struct {
	GRPCMicroserviceAddress string
}

func Create(conf *Config) (*UserService, error) {
	userMicroserviceConnect, err := grpc.Dial(conf.GRPCMicroserviceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	userServiceClient := proto.NewUserServiceClient(userMicroserviceConnect)

	userService := &UserService{
		userServiceClient: userServiceClient,
		ctx:               context.Background(),
	}

	return userService, nil
}
