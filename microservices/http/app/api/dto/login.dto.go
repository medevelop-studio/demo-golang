package dto

import (
	"http/app/api/codes"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id    string         `json:"id"`
	Login string         `json:"login"`
	Role  codes.UserRole `json:"role"`
	Token string         `json:"token"`
}
