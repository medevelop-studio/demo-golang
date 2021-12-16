package user_service

import user_domain "user/app/domain/user"

func (service *UserService) Create(data *user_domain.CreateUserDto) (*user_domain.User, error) {
	user, err := user_domain.Create(data)

	if err != nil {
		return nil, err
	}

	if err := service.store.SaveUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
