package codes

import (
	"errors"
	"fmt"
)

type ErrorStatusCode byte

const (
	ErrJSONFormat ErrorStatusCode = iota + 1
	ErrJSONValidation

	ErrUnknown
	ErrInternalServer

	ErrLoginInvalidCreds
)

func ErrorStatusCodeToError(code ErrorStatusCode) error {
	return errors.New(fmt.Sprint(code))
}
