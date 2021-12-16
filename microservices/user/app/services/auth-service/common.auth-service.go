package auth_service

import user_service "user/app/services/user-service"

type AuthService struct {
	config      *Config
	userService *user_service.UserService
}

type Config struct {
	TokenSecret   string
	TokenLifeTime int64
}

func Create(config *Config, userService *user_service.UserService) (*AuthService, error) {
	authService := &AuthService{
		config:      config,
		userService: userService,
	}

	return authService, nil
}
