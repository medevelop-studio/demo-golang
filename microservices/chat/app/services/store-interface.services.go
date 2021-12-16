package services

import message_service "chat/app/services/message-service"

type Store interface {
	message_service.Store

	Close()
}
