package driver_interfaces

import user_domain "user/app/domain/user"

type UserDb interface {
	SaveUser(user *user_domain.User) error
	GetUserById(id string) (*user_domain.User, error)
	GetUserByLogin(login string) (*user_domain.User, error)
}
