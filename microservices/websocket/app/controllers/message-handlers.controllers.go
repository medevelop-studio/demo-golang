package controllers

import (
	"log"
	"websocket/app/api/dto"
)

func (controllers *Controllers) handlerMessageRequest(
	data []byte, id string, send func(data []byte) error,
) error {
	var req dto.ChatMessageWTNRequest
	if err := controllers.bind(data, &req); err != nil {
		log.Println(err)
	}

	req.Data.UserId = id

	return controllers.natsConnections.SendMessageRequest(req)
}

func (controllers *Controllers) handlerMessageResponse(
	data []byte, id string, send func(data []byte) error,
) error {
	var req dto.ChatMessageWTNRequest
	if err := controllers.bind(data, &req); err != nil {
		log.Println(err)
	}

	controllers.wsConnections.SendToAuthorizedConnections(req)

	return nil
}
