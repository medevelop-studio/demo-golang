package router

func (router *Router) NewMessage(data []byte, connId string) error {
	code, err := router.getMessageCode(data)
	if err != nil {
		return router.handleError(data, connId, err)
	}

	handler, ok := router.handlers[code]
	if ok {
		return handler(data, connId)
	}

	if err := router.handleDefault(data, connId); err != nil {
		return router.handleError(data, connId, err)
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

func (router *Router) handleDefault(data []byte, connId string) error {
	if router.defaultHandler == nil {
		return ErrDefaultHandlerNotSet
	}

	return router.defaultHandler(data, connId)
}
