package message_service

import (
	"chat/app/api/dto"
	message_domain "chat/app/domain/message-domain"
)

func (service *MessageService) NewMessageToChat(data *dto.ChatMessageData) (*dto.MessageToChat, error) {
	user, err := service.userService.GetUserById(data.UserId)
	if err != nil {
		return nil, err
	}

	message, err := message_domain.Create(&message_domain.CreateMessageDto{
		UserId:  data.UserId,
		Content: data.Content,
	})
	if err != nil {
		return nil, err
	}

	if err := service.store.SaveMessage(message); err != nil {
		return nil, err
	}

	result := &dto.MessageToChat{
		Id:       message.Id,
		UserId:   message.UserId,
		Content:  message.Content,
		Date:     message.Date,
		Username: user.Login,
	}

	return result, nil
}
