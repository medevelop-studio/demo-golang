package services

import (
	message_service "chat/app/services/message-service"
	user_service "chat/app/services/user-service"
)

type Services struct {
	MessageService *message_service.MessageService
	UserService    *user_service.UserService
}

type Config struct {
	GRPCUserMicroserviceAddress string
}

func Create(store Store, conf *Config) (*Services, error) {
	userService, err := user_service.Create(&user_service.Config{
		GRPCMicroserviceAddress: conf.GRPCUserMicroserviceAddress,
	})
	if err != nil {
		return nil, err
	}

	messageService, err := message_service.Create(store, userService)
	if err != nil {
		return nil, err
	}

	services := &Services{
		UserService:    userService,
		MessageService: messageService,
	}

	return services, nil
}
