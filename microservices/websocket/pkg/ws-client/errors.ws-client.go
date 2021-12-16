package ws_client

import (
	"errors"
)

var (
	ErrClientCloseConnect = errors.New("client socket: error connection is close")

	ErrClientHandlerMessageNotSet = errors.New("client socket: error handler message not set")
)
