package services

import user_domain "user/app/domain/user"

func (services *Services) createDemoData() error {
	_, err := services.UserService.Create(&user_domain.CreateUserDto{
		Login:    "admin",
		Password: "admin",
		Role:     user_domain.ROLE_ADMIN,
	})
	if err != nil {
		return err
	}

	_, err = services.UserService.Create(&user_domain.CreateUserDto{
		Login:    "user1",
		Password: "1111",
		Role:     user_domain.ROLE_COMMON,
	})
	if err != nil {
		return err
	}

	_, err = services.UserService.Create(&user_domain.CreateUserDto{
		Login:    "user2",
		Password: "1111",
		Role:     user_domain.ROLE_COMMON,
	})
	if err != nil {
		return err
	}

	return nil
}
