package user_domain

import "errors"

var (
	ErrCreateLoginOrPasswordEmpty = errors.New("field login or password is empty")
)
