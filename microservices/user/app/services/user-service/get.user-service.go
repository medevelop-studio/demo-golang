package user_service

import (
	user_domain "user/app/domain/user"
)

func (service *UserService) GetUserById(id string) (*user_domain.User, error) {
	user, err := service.store.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) GetUserByLogin(login string) (*user_domain.User, error) {
	user, err := service.store.GetUserByLogin(login)

	if err != nil {
		return nil, err
	}

	return user, nil
}
