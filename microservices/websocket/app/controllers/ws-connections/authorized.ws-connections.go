package ws_connections

import (
	"time"
	ws_client "websocket/pkg/ws-client"
)

func (conns *WsConnections) AuthorizeConnection(connId string, userId string, message interface{}) error {
	clientObj, ok := conns.unauthorizedClients.LoadAndDelete(connId)
	if !ok {
		return ErrClientNotFound
	}
	client := clientObj.(*ws_client.Client)

	if client.Status == ws_client.CLIENT_STATUS_STOP {
		return ErrClientCloseConnect
	}

	if previesConnect, ok := conns.authorizedClients.LoadAndDelete(userId); ok {
		previesConnect.(*ws_client.Client).Stop()
	}

	client.Status = ws_client.CLIENT_STATUS_READY
	client.ConnId = userId
	client.SetHandlerMessage(conns.authorizedRouter.NewMessage)

	conns.authorizedClients.Store(userId, client)

	if message != nil {
		if err := client.SendMessage(message); err != nil {
			return err
		}
	}

	return nil
}

func (conns *WsConnections) DeleteAuthorizedConnection(connId string, message interface{}) error {
	clientObj, ok := conns.authorizedClients.LoadAndDelete(connId)
	if !ok {
		return ErrClientNotFound
	}
	client := clientObj.(*ws_client.Client)

	if message != nil {
		if err := client.SendMessage(message); err != nil {
			return err
		}
	}

	time.Sleep(1 * time.Second)

	client.Stop()

	return nil
}

func (conns *WsConnections) SendToAuthorizedConnections(data interface{}) {
	conns.authorizedClients.Range(func(key, value interface{}) bool {
		client := value.(*ws_client.Client)
		client.SendMessage(data)
		return true
	})
}
