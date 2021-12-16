package controllers

import (
	"log"
	"websocket/app/api/codes"
	"websocket/app/api/dto"

	"github.com/gorilla/websocket"
)

func (controllers *Controllers) handlerNewConnection(conn *websocket.Conn) error {
	_, err := controllers.wsConnections.AddUnauthorizedConnect(conn)

	return err
}

func (controllers *Controllers) handlerLoginUserRequest(
	data []byte, id string, send func(data []byte) error,
) error {
	var req dto.LoginUTWRequest
	if err := controllers.bind(data, &req); err != nil {
		log.Println(err)
	}

	natsReq := dto.LoginWTNRequest{
		RequestGeneral: dto.RequestGeneral{
			Type: codes.LoginRequest,
		},
		Login:        req.Data.Login,
		Password:     req.Data.Password,
		ConnectionId: id,
	}

	return controllers.natsConnections.SendLoginRequest(natsReq)
}

func (controllers *Controllers) handlerLoginUserResponse(
	data []byte, id string, send func(data []byte) error,
) error {
	var req dto.LoginNTWResponse
	if err := controllers.bind(data, &req); err != nil {
		log.Println(err)
	}

	if req.Error != nil {
		return controllers.wsConnections.DeleteUnauthorizedUser(req.ConnectionId, req)
	}

	wsRes := dto.LoginWTUResponse{
		ResponseGeneral: dto.ResponseGeneral{
			Type: codes.LoginResponse,
		},
		User: req.Data,
	}

	err := controllers.wsConnections.AuthorizeConnection(req.ConnectionId, req.Data.Id, wsRes)
	if err != nil {
		return controllers.wsConnections.DeleteUnauthorizedUser(req.ConnectionId, nil)
	}

	return nil
}
