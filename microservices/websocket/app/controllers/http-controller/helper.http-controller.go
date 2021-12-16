package http_controller

import (
	"encoding/json"
	"net/http"
)

func (controller *HttpController) sendErrorBadRequestResponse(w http.ResponseWriter, err error) error {
	http.Error(w, err.Error(), http.StatusBadRequest)
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
