package nats_controller

import (
	"chat/app/api/codes"

	"github.com/nats-io/nats.go"
)

type Router struct {
	handlers       map[codes.MessageTypeCode]HandlerFunc
	defaultHandler HandlerFunc
	parserHandler  func(message []byte) (codes.MessageTypeCode, error)
	errorHandler   func(*nats.Msg, string, error)
}

func CreateRouter() *Router {
	return &Router{
		handlers: make(map[codes.MessageTypeCode]HandlerFunc),
	}
}

func (router *Router) SetParser(handler func(message []byte) (codes.MessageTypeCode, error)) {
	router.parserHandler = handler
}

func (router *Router) SetDefaultHandler(handler HandlerFunc) {
	router.defaultHandler = handler
}

func (router *Router) SetErrorHandler(handler func(*nats.Msg, string, error)) {
	router.errorHandler = handler
}

func (router *Router) AddHandler(code codes.MessageTypeCode, handler HandlerFunc) {
	router.handlers[code] = handler
}

func (router *Router) NewMessage(message *nats.Msg, subjectId string) error {
	code, err := router.getMessageCode(message.Data)
	if err != nil {
		return router.handleError(message, subjectId, err)
	}

	handler, ok := router.handlers[code]
	if ok {
		return handler(message, subjectId)
	}

	if err := router.handleDefault(message, subjectId); err != nil {
		return router.handleError(message, subjectId, err)
	}

	return nil
}

func (router *Router) getMessageCode(data []byte) (codes.MessageTypeCode, error) {
	if router.parserHandler == nil {
		return 0, ErrRouterParserNotSet
	}

	code, err := router.parserHandler(data)
	if err != nil {
		return 0, err
	}

	return code, nil
}

func (router *Router) handleError(message *nats.Msg, subjectId string, err error) error {
	if router.errorHandler == nil {
		return ErrRouterErrorHandlerNotSet
	}

	router.errorHandler(message, subjectId, err)
	return nil
}

func (router *Router) handleDefault(message *nats.Msg, subjectId string) error {
	if router.defaultHandler == nil {
		return ErrRouterDefaultHandlerNotSet
	}

	return router.defaultHandler(message, subjectId)
}
