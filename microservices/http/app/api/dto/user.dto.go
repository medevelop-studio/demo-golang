package dto

import "http/app/api/codes"

type User struct {
	Id    string         `json:"id"`
	Login string         `json:"login"`
	Role  codes.UserRole `json:"role"`
}

type CreateUserRequest struct {
	Login    string         `json:"login" validate:"required,max=32"`
	Password string         `json:"password" validate:"required,min=5,max=35"`
	Role     codes.UserRole `json:"role" validate:"number,max=2"`
}
