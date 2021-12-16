package message_service

import user_service "chat/app/services/user-service"

type MessageService struct {
	store       Store
	userService *user_service.UserService
}

func Create(store Store, userService *user_service.UserService) (*MessageService, error) {
	messageService := &MessageService{
		store:       store,
		userService: userService,
	}

	return messageService, nil
}
