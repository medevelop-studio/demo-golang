package mock_db

import (
	message_domain "chat/app/domain/message-domain"
)

type MessageDB struct {
	messages map[string]*message_domain.Message
}

func Create() (*MessageDB, error) {
	return &MessageDB{
		messages: make(map[string]*message_domain.Message),
	}, nil
}

func (db *MessageDB) SaveMessage(message *message_domain.Message) error {
	db.messages[message.Id] = message

	return nil
}

func (db *MessageDB) GetMessageById(id string) (*message_domain.Message, error) {
	user, ok := db.messages[id]

	if !ok {
		return nil, nil
	}

	return user, nil
}
