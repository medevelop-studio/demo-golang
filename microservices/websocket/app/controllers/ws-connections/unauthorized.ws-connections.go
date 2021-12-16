package ws_connections

import (
	"crypto/rand"
	"fmt"
	"time"
	ws_client "websocket/pkg/ws-client"

	"github.com/gorilla/websocket"
)

func (conns *WsConnections) AddUnauthorizedConnect(conn *websocket.Conn) (*ws_client.Client, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	connId := fmt.Sprintf("%x", buf)

	client := ws_client.CreateClient(connId, conn)
	client.SetHandlerMessage(conns.unauthorizedRouter.NewMessage)
	client.SetHandlerCloseConn(conns.closeConnHandler)
	if err := client.Run(); err != nil {
		return nil, err
	}

	conns.unauthorizedClients.Store(connId, client)

	return client, nil
}

func (conns *WsConnections) DeleteUnauthorizedUser(connId string, message interface{}) error {
	clientObj, ok := conns.unauthorizedClients.LoadAndDelete(connId)
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
