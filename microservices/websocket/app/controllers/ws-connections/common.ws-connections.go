package ws_connections

import (
	"sync"
	"websocket/pkg/router"
)

type WsConnections struct {
	unauthorizedRouter  *router.Router
	authorizedRouter    *router.Router
	unauthorizedClients sync.Map
	authorizedClients   sync.Map
	closeConnHandler    func(data []byte, connId string, send func(data []byte) error) error
}

func Create() (*WsConnections, error) {
	wsConnections := &WsConnections{
		unauthorizedClients: sync.Map{},
		authorizedClients:   sync.Map{},
	}

	return wsConnections, nil
}

func (conns *WsConnections) SetUnauthorizedRouter(router *router.Router) {
	conns.unauthorizedRouter = router
}

func (conns *WsConnections) SetAuthorizedRouter(router *router.Router) {
	conns.authorizedRouter = router
}
