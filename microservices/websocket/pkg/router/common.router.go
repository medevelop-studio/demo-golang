package router

type SendFunction func(data []byte) error
type MessageHandlerFunc func(data []byte, connId string, send func(data []byte) error) error
type ErrorHandlerFunc func(data []byte, connId string, err error)
type ParserFunc func(data []byte) (interface{}, error)

type Router struct {
	handlers       map[interface{}]MessageHandlerFunc
	defaultHandler MessageHandlerFunc
	parserHandler  ParserFunc
	errorHandler   ErrorHandlerFunc
}

func CreateRouter() *Router {
	return &Router{
		handlers: make(map[interface{}]MessageHandlerFunc),
	}
}

func (router *Router) SetMessageParser(handler ParserFunc) {
	router.parserHandler = handler
}

func (router *Router) SetDefaultHandler(handler MessageHandlerFunc) {
	router.defaultHandler = handler
}

func (router *Router) SetErrorHandler(handler ErrorHandlerFunc) {
	router.errorHandler = handler
}

func (router *Router) AddHandler(code interface{}, handler MessageHandlerFunc) {
	router.handlers[code] = handler
}

func (router *Router) NewMessage(message []byte, connId string, send func(data []byte) error) error {
	code, err := router.getMessageCode(message)
	if err != nil {
		return router.handleError(message, connId, err)
	}

	handler, ok := router.handlers[code]
	if ok {
		return handler(message, connId, send)
	}

	if err := router.handleDefault(message, connId, send); err != nil {
		return router.handleError(message, connId, err)
	}

	return nil
}

func (router *Router) getMessageCode(data []byte) (interface{}, error) {
	if router.parserHandler == nil {
		return 0, ErrParserNotSet
	}

	code, err := router.parserHandler(data)
	if err != nil {
		return 0, err
	}

	return code, nil
}

func (router *Router) handleError(data []byte, connId string, err error) error {
	if router.errorHandler == nil {
		return ErrErrorHandlerNotSet
	}

	router.errorHandler(data, connId, err)
	return nil
}

func (router *Router) handleDefault(message []byte, connId string, send SendFunction) error {
	if router.defaultHandler == nil {
		return ErrDefaultHandlerNotSet
	}

	return router.defaultHandler(message, connId, send)
}
