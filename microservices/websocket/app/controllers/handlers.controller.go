package controllers

import (
	"encoding/json"
	"websocket/app/api/codes"
	"websocket/app/api/dto"
	"websocket/pkg/router"
)

func (controllers *Controllers) setHandlers() {
	controllers.httpController.SetNewConnectionHandler(controllers.handlerNewConnection)

	unauthorizedUserWsRouter := router.CreateRouter()
	unauthorizedUserWsRouter.SetMessageParser(controllers.parseMesssageCode)
	unauthorizedUserWsRouter.AddHandler(codes.LoginRequest, controllers.handlerLoginUserRequest)
	controllers.wsConnections.SetUnauthorizedRouter(unauthorizedUserWsRouter)

	authorizedUserWsRouter := router.CreateRouter()
	authorizedUserWsRouter.SetMessageParser(controllers.parseMesssageCode)
	authorizedUserWsRouter.AddHandler(codes.ChatMessageRequest, controllers.handlerMessageRequest)
	controllers.wsConnections.SetAuthorizedRouter(authorizedUserWsRouter)

	serverNatsRouter := router.CreateRouter()
	serverNatsRouter.SetMessageParser(controllers.parseMesssageCode)
	serverNatsRouter.AddHandler(codes.LoginResponse, controllers.handlerLoginUserResponse)
	controllers.natsConnections.SetServerRouter(serverNatsRouter)

	messageNatsRouter := router.CreateRouter()
	messageNatsRouter.SetMessageParser(controllers.parseMesssageCode)
	messageNatsRouter.AddHandler(codes.ChatMessageResponse, controllers.handlerMessageResponse)
	controllers.natsConnections.SetMessageRouter(messageNatsRouter)
}

func (controllers *Controllers) parseMesssageCode(data []byte) (interface{}, error) {
	var message dto.RequestGeneral
	if err := json.Unmarshal(data, &message); err != nil {
		return 0, err
	}

	return message.Type, nil
}

// func (controllers *Controllers) parseNatsResponse(data []byte) (interface{}, error) {
// 	var message dto.ResponseGeneral
// 	if err := json.Unmarshal(data, &message); err != nil {
// 		return 0, err
// 	}

// 	return message.Type, nil
// }
