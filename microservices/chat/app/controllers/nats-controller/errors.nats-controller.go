package nats_controller

import "errors"

var (
	ErrRouterParserNotSet         = errors.New("error parser handler not set")
	ErrRouterErrorHandlerNotSet   = errors.New("error error handler not set")
	ErrRouterDefaultHandlerNotSet = errors.New("error default handler not set")
)
