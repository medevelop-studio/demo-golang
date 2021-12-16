package dto

type LoginUTWRequest struct {
	RequestGeneral
	Data struct {
		Login    string `json:"login"  validate:"required"`
		Password string `json:"password"  validate:"required"`
	} `json:"data"  validate:"required"`
}

type LoginWTUResponse struct {
	ResponseGeneral
	User *UserLoginData `json:"user,omitempty"`
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
	Data         *UserLoginData `json:"user"`
}

type UserLoginData struct {
	Id    string   `json:"id,omitempty"`
	Login string   `json:"login,omitempty"`
	Role  UserRole `json:"role,omitempty"`
}
