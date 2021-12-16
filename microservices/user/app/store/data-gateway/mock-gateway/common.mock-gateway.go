package mock_gateway

import (
	driver_interfaces "user/app/store/data-gateway/driver-interfaces"
)

type MockGateway struct {
	UserDb driver_interfaces.UserDb
}

func (gateway *MockGateway) Close() {}
