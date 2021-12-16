package ws_connections

import "errors"

var (
	ErrClientNotFound     = errors.New("connections: client not found")
	ErrClientCloseConnect = errors.New("connections: error client connection is close")
)
