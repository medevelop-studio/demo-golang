package services

import (
	auth_service "http/app/services/auth-service"
	user_service "http/app/services/user-service"

	"google.golang.org/grpc"
)

type Services struct {
	AuthService *auth_service.AuthService
	UserService *user_service.UserService
}

type Config struct {
	GRPCUserMicroserviceAddress string
}

func Create(conf *Config) (*Services, error) {
	userMicroserviceConnect, err := grpc.Dial(conf.GRPCUserMicroserviceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	authService, err := auth_service.Create(userMicroserviceConnect)
	if err != nil {
		return nil, err
	}

	userService, err := user_service.Create(userMicroserviceConnect)
	if err != nil {
		return nil, err
	}

	services := &Services{
		AuthService: authService,
		UserService: userService,
	}

	return services, nil
}
