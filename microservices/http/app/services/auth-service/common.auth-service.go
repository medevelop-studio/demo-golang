package auth_service

import (
	"context"
	"http/proto"

	"google.golang.org/grpc"
)

type AuthService struct {
	authServiceConnection *grpc.ClientConn
	authServiceClient     proto.AuthServiceClient
	ctx                   context.Context
}

func Create(userMicroserviceConn *grpc.ClientConn) (*AuthService, error) {
	authServiceClient := proto.NewAuthServiceClient(userMicroserviceConn)

	authService := &AuthService{
		authServiceConnection: userMicroserviceConn,
		authServiceClient:     authServiceClient,
		ctx:                   context.Background(),
	}

	return authService, nil
}
