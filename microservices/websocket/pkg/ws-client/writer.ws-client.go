package ws_client

import (
	"time"

	"github.com/gorilla/websocket"
)

func (client *Client) writer() {
	ticker := time.NewTicker(DEFAULT_PING_PERIOD)
	defer func() {
		ticker.Stop()
		if recoveryMessage := recover(); recoveryMessage != nil {
			client.handlerPanic(recoveryMessage)
		}
	}()

	for {
		select {
		case <-client.ctx.Done():
			return
		case message, ok := <-client.sendToClient:
			if !ok {
				return
			}
			client.conn.SetWriteDeadline(time.Now().Add(DEFAULT_WRITE_WAIT))

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				client.handlerError(err, false)
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				client.handlerError(err, false)
				return
			}

		case message, ok := <-client.sendToClientBuffered:
			if !ok {
				return
			}
			client.conn.SetWriteDeadline(time.Now().Add(DEFAULT_WRITE_WAIT))

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				client.handlerError(err, false)
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				client.handlerError(err, false)
				return
			}

			if len(client.sendToClientBuffered) > 3 {
				<-client.sendToClientBuffered
			}

		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(DEFAULT_WRITE_WAIT))
			err := client.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				client.handlerError(err, false)
				return
			}
		}
	}
}
