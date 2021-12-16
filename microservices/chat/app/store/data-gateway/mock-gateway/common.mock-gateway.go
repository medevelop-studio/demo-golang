package mock_gateway

import (
	driver_interfaces "chat/app/store/data-gateway/driver-interfaces"
)

type MockGateway struct {
	MessageDb driver_interfaces.MessageDb
}

func (gateway *MockGateway) Close() {}
