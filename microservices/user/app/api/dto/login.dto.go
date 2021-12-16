package dto

import user_domain "user/app/domain/user"

type AdminLoginData struct {
	Id    string               `json:"id"`
	Login string               `json:"login"`
	Role  user_domain.UserRole `json:"role"`
	Token string               `json:"token"`
}

type LoginWTNRequest struct {
	RequestGeneral
	Login        string `json:"login"`
	Password     string `json:"password"`
	ConnectionId string `json:"connectionId"`
	ServerId     string `json:"serverId"`
}

type LoginNTWResponse struct {
	ResponseGeneral
	ConnectionId string         `json:"connectionId"`
	Data         *UserLoginData `json:"user,omitempty"`
}

type UserLoginData struct {
	Id    string               `json:"id,omitempty"`
	Login string               `json:"login,omitempty"`
	Role  user_domain.UserRole `json:"role,omitempty"`
}
