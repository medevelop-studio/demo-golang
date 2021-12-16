package store

import (
	"chat/app/services"
	mock_gateway "chat/app/store/data-gateway/mock-gateway"
	mock_db "chat/app/store/drivers/mock-db"
)

func CreateMockStore() (services.Store, error) {
	messageDb, err := mock_db.Create()
	if err != nil {
		return nil, err
	}

	return &mock_gateway.MockGateway{
		MessageDb: messageDb,
	}, nil
}
