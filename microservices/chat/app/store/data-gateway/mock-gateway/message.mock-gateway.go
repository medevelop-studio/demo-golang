package mock_gateway

import message_domain "chat/app/domain/message-domain"

func (gateway *MockGateway) SaveMessage(message *message_domain.Message) error {
	return gateway.MessageDb.SaveMessage(message)
}

func (gateway *MockGateway) GetMessageById(id string) (*message_domain.Message, error) {
	user, err := gateway.MessageDb.GetMessageById(id)

	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, ErrMessageNotFound
	}

	return user, nil
}
