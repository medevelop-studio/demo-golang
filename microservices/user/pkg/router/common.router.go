package router

type HandlerFunc func([]byte, string) error

type Router struct {
	handlers       map[interface{}]HandlerFunc
	defaultHandler HandlerFunc
	parserHandler  func([]byte) (interface{}, error)
	errorHandler   func([]byte, string, error)
}

func CreateRouter() *Router {
	return &Router{
		handlers: make(map[interface{}]HandlerFunc),
	}
}

func (router *Router) SetParser(handler func([]byte) (interface{}, error)) {
	router.parserHandler = handler
}

func (router *Router) SetDefaultHandler(handler HandlerFunc) {
	router.defaultHandler = handler
}

func (router *Router) SetErrorHandler(handler func([]byte, string, error)) {
	router.errorHandler = handler
}

func (router *Router) AddHandler(code interface{}, handler HandlerFunc) {
	router.handlers[code] = handler
}
