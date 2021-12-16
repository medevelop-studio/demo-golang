package store

import (
	"user/app/services"
	mock_gateway "user/app/store/data-gateway/mock-gateway"
	mock_db "user/app/store/drivers/mock-db"
)

func CreateMockStore() (services.Store, error) {
	userDb, err := mock_db.Create()
	if err != nil {
		return nil, err
	}

	return &mock_gateway.MockGateway{
		UserDb: userDb,
	}, nil
}
