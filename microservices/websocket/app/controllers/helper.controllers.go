package controllers

import (
	"encoding/json"
)

func (controllers *Controllers) bind(data []byte, value interface{}) error {
	return json.Unmarshal(data, value)
}

// func (controllers *Controllers) validate(value interface{}) error {
// 	return validator.New().Struct(value)
// }

// func (controllers *Controllers)
