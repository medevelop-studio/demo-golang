package user_service

type UserService struct {
	store Store
}

func Create(store Store) (*UserService, error) {
	userService := &UserService{
		store: store,
	}

	return userService, nil
}
