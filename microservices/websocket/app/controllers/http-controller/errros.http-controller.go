package http_controller

import "errors"

var (
	ErrMethodNotSupport  = errors.New("this method not supported")
	ErrServerUnavailable = errors.New("server unavailable")
)
