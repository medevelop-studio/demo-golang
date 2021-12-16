package auth_service

import "errors"

var (
	ErrUnexpectedMethod = errors.New("unexpected signing method")
)
