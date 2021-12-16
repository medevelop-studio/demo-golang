package auth_service

import (
	"user/app/api/codes"
	"user/app/api/dto"
	user_domain "user/app/domain/user"
)

func (service *AuthService) LoginUser(login, password string) (*dto.UserLoginData, error) {
	user, err := service.commonLogin(login, password)
	if err != nil {
		return nil, err
	}

	loginData := &dto.UserLoginData{
		Id:    user.Id,
		Login: user.Login,
		Role:  user.Role,
	}

	return loginData, nil
}

func (service *AuthService) LoginAdmin(login, password string) (*dto.AdminLoginData, error) {
	user, err := service.commonLogin(login, password)
	if err != nil {
		return nil, err
	}

	if !user.IsAdmin() {
		return nil, codes.ErrorStatusCodeToError(codes.ErrLoginInvalidCreds)
	}

	token, err := service.createJWTToken(user)
	if err != nil {
		return nil, err
	}

	loginData := &dto.AdminLoginData{
		Id:    user.Id,
		Login: user.Login,
		Role:  user.Role,
		Token: token,
	}

	return loginData, nil
}

func (service *AuthService) commonLogin(login, password string) (*user_domain.User, error) {
	user, err := service.userService.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}

	if !user.IsPasswordCorrect(password) {
		return nil, codes.ErrorStatusCodeToError(codes.ErrLoginInvalidCreds)
	}

	return user, nil
}
