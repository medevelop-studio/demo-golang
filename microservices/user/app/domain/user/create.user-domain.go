package user_domain

import "github.com/google/uuid"

const DEFAULT_USER_ROLE UserRole = ROLE_COMMON

func Create(data *CreateUserDto) (*User, error) {
	if data.Login == "" || data.Password == "" {
		return nil, ErrCreateLoginOrPasswordEmpty
	}

	user := &User{
		Id:    uuid.NewString(),
		Login: data.Login,
	}

	if err := user.SetPassword(data.Password); err != nil {
		return nil, err
	}

	if data.Role != 0 {
		user.Role = data.Role
	} else {
		user.Role = DEFAULT_USER_ROLE
	}

	return user, nil
}
