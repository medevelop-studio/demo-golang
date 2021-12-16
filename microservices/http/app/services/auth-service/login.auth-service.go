package auth_service

import (
	"http/app/api/codes"
	"http/app/api/dto"
	"http/proto"
)

func (service *AuthService) Login(data *dto.LoginRequest) (*dto.LoginResponse, error) {
	protoReq := &proto.LoginRequest{
		Login:    data.Login,
		Password: data.Password,
	}

	res, err := service.authServiceClient.Login(service.ctx, protoReq)
	if err != nil {
		return nil, err
	}

	result := &dto.LoginResponse{
		Id:    res.User.Id,
		Login: res.User.Login,
		Role:  codes.UserRole(res.User.Role),
		Token: res.Token,
	}

	return result, err
}
