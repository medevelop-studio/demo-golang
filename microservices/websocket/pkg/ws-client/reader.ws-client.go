package ws_client

import (
	"github.com/gorilla/websocket"
)

func (client *Client) reader() {
	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			client.handlerPanic(recoveryMessage)
		}
	}()

	for {
		select {
		case <-client.ctx.Done():
			return
		default:
			message, err := readMessage(client.conn)
			if err != nil {
				if client.Status != CLIENT_STATUS_STOP {
					client.handlerError(err, true)
				}
				return
			}
			client.handlerMessage(message)
		}
	}
}

func readMessage(conn *websocket.Conn) ([]byte, error) {
	_, message, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return message, nil
}
