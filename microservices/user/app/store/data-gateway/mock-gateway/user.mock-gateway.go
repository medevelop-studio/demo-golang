package mock_gateway

import user_domain "user/app/domain/user"

func (gateway *MockGateway) SaveUser(user *user_domain.User) error {
	return gateway.UserDb.SaveUser(user)
}

func (gateway *MockGateway) GetUserById(id string) (*user_domain.User, error) {
	user, err := gateway.UserDb.GetUserById(id)

	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (gateway *MockGateway) GetUserByLogin(login string) (*user_domain.User, error) {
	user, err := gateway.UserDb.GetUserByLogin(login)

	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
