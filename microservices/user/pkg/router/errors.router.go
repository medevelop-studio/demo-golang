package router

import "errors"

var (
	ErrParserNotSet         = errors.New("error parser handler not set")
	ErrErrorHandlerNotSet   = errors.New("error error handler not set")
	ErrDefaultHandlerNotSet = errors.New("error default handler not set")
)
