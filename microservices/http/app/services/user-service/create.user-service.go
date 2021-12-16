package user_service

import (
	"http/app/api/codes"
	"http/app/api/dto"
	"http/proto"
)

func (service *UserService) CreateUser(data *dto.CreateUserRequest) (*dto.User, error) {
	protoReq := &proto.CreateUserRequest{
		Login:    data.Login,
		Password: data.Password,
		Role:     proto.UserRoles(data.Role),
	}

	res, err := service.userServiceClient.CreateUser(service.ctx, protoReq)
	if err != nil {
		return nil, err
	}

	result := &dto.User{
		Id:    res.Id,
		Login: res.Login,
		Role:  codes.UserRole(res.Role),
	}

	return result, err
}
