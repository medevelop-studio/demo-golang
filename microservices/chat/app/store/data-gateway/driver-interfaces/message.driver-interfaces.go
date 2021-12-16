package driver_interfaces

import message_domain "chat/app/domain/message-domain"

type MessageDb interface {
	SaveMessage(message *message_domain.Message) error
	GetMessageById(id string) (*message_domain.Message, error)
}
