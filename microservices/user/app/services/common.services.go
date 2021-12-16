package services

import (
	auth_service "user/app/services/auth-service"
	user_service "user/app/services/user-service"
)

type Services struct {
	UserService *user_service.UserService
	AuthService *auth_service.AuthService
}

type Config struct {
	AuthTokenSecret   string
	AuthTokenLifeTime int64
}

func Create(store Store, config *Config) (*Services, error) {
	userService, err := user_service.Create(store)
	if err != nil {
		return nil, err
	}

	authConfig := &auth_service.Config{
		TokenSecret:   config.AuthTokenSecret,
		TokenLifeTime: config.AuthTokenLifeTime,
	}
	authService, err := auth_service.Create(authConfig, userService)
	if err != nil {
		return nil, err
	}

	services := &Services{
		UserService: userService,
		AuthService: authService,
	}

	if err := services.createDemoData(); err != nil {
		return nil, err
	}

	return services, nil
}
