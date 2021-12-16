package http_controller

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (controller *HttpController) addNewConnectionHandler() {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  0,
		WriteBufferSize: 0,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			controller.sendErrorBadRequestResponse(w, ErrMethodNotSupport)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			conn.WriteJSON(map[string]interface{}{
				"statusCode": http.StatusInternalServerError,
				"error":      err.Error(),
			})
			conn.Close()
			return
		}

		if controller.newConnHandler == nil {
			conn.WriteJSON(map[string]interface{}{
				"statusCode": http.StatusInternalServerError,
				"error":      ErrServerUnavailable.Error(),
			})
			conn.Close()
			return
		}

		if err := controller.newConnHandler(conn); err != nil {
			conn.WriteJSON(map[string]interface{}{
				"statusCode": http.StatusInternalServerError,
				"error":      ErrServerUnavailable.Error(),
			})
			conn.Close()
			return
		}
	})
}
