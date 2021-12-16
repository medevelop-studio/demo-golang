package services

import user_domain "user/app/domain/user"

type Store interface {
	SaveUser(user *user_domain.User) error
	GetUserById(id string) (*user_domain.User, error)
	GetUserByLogin(login string) (*user_domain.User, error)

	Close()
}
