package dto

import (
	"chat/app/api/codes"
	"strconv"
)

type RequestGeneral struct {
	Type codes.MessageTypeCode `json:"type"`
}

type ResponseGeneral struct {
	Type    codes.MessageTypeCode `json:"type"`
	Error   *Error                `json:"error,omitempty"`
	Data    interface{}           `json:"data,omitempty"`
	Request interface{}           `json:"request,omitempty"`
}

func CreateResponse(
	resType codes.MessageTypeCode,
) ResponseGeneral {
	return ResponseGeneral{
		Type: resType,
	}
}

type Error struct {
	Code    codes.ErrorStatusCode `json:"code"`
	Message string                `json:"message,omitempty"`
}

func ConvertErrToError(err error) *Error {
	if errCode, parseErr := strconv.ParseUint(err.Error(), 10, 8); parseErr == nil {
		return &Error{
			Code: codes.ErrorStatusCode(errCode),
		}
	} else {
		return &Error{
			Code:    codes.ErrUnknown,
			Message: err.Error(),
		}
	}
}
