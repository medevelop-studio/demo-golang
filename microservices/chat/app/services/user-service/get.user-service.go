package user_service

import (
	user_domain "chat/app/domain/user-domain"
	"chat/proto"
)

func (service *UserService) GetUserById(id string) (*user_domain.User, error) {
	res, err := service.userServiceClient.GetUserById(service.ctx, &proto.GetUserByIdRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	user := &user_domain.User{
		Id:    res.Id,
		Login: res.Login,
		Role:  user_domain.UserRole(res.Role),
	}

	return user, nil
}
