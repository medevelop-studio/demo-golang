package ws_client

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type RequestCtx struct {
	ConnId      string
	ConnStatus  ClientStatusCode
	Data        *[]byte
	Error       error
	SendMessage func(interface{}) error
}

func (ctx *RequestCtx) Bind(value interface{}) error {
	return json.Unmarshal(*ctx.Data, value)
}
func (ctx *RequestCtx) Validate(value interface{}) error {
	return validator.New().Struct(value)
}
