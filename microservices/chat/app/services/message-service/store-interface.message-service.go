package message_service

import message_domain "chat/app/domain/message-domain"

type Store interface {
	SaveMessage(message *message_domain.Message) error
	GetMessageById(id string) (*message_domain.Message, error)
}
