package http_controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (controller *HttpController) Bind(body io.ReadCloser, obj interface{}) error {
	return json.NewDecoder(body).Decode(obj)
}

func (controller *HttpController) ValidateOrSendError(w http.ResponseWriter, value interface{}) error {
	err := validator.New().Struct(value)

	if err == nil {
		return nil
	}

	controller.sendErrorBadRequestResponse(w, err)

	return err
}

func (controller *HttpController) sendSuccessResponse(w http.ResponseWriter, data interface{}) error {
	return controller.sendResponse(w, http.StatusOK, data)
}

func (controller *HttpController) sendErrorResponse(w http.ResponseWriter, err error) error {
	return controller.sendResponse(w, http.StatusInternalServerError, map[string]interface{}{
		"statusCode": http.StatusInternalServerError,
		"error":      err.Error(),
	})
}

func (controller *HttpController) sendErrorBadRequestResponse(w http.ResponseWriter, err error) error {
	return controller.sendResponse(w, http.StatusBadRequest, map[string]interface{}{
		"statusCode": http.StatusBadRequest,
		"error":      err.Error(),
	})
}

func (controller *HttpController) sendResponse(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}

	return nil
}
