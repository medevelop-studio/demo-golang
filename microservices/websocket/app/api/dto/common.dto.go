package dto

import (
	"websocket/app/api/codes"
)

type Error struct {
	Code    codes.ErrorStatusCode `json:"code"`
	Message string                `json:"message,omitempty"`
}

type RequestGeneral struct {
	Type codes.MessageTypeCode `json:"type"`
}

type ResponseGeneral struct {
	Type    codes.MessageTypeCode `json:"type"`
	Error   *Error                `json:"error,omitempty"`
	Data    interface{}           `json:"data,omitempty"`
	Request interface{}           `json:"request,omitempty"`
}
