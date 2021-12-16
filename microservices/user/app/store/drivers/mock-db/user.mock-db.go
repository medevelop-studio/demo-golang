package mock_db

import user_domain "user/app/domain/user"

type UserDB struct {
	Users map[string]*user_domain.User
}

func Create() (*UserDB, error) {
	return &UserDB{
		Users: make(map[string]*user_domain.User),
	}, nil
}

func (db *UserDB) SaveUser(user *user_domain.User) error {
	db.Users[user.Id] = user

	return nil
}

func (db *UserDB) GetUserById(id string) (*user_domain.User, error) {
	user, ok := db.Users[id]

	if !ok {
		return nil, nil
	}

	return user, nil
}

func (db *UserDB) GetUserByLogin(login string) (*user_domain.User, error) {
	for _, user := range db.Users {
		if user.Login == login {
			return user, nil
		}
	}

	return nil, nil
}
