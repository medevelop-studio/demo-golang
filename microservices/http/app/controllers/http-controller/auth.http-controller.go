package http_controller

import (
	"http/app/api/dto"
	"net/http"
)

func (controller *HttpController) addAuthHandlers() {
	controller.router.HandleFunc("/login", controller.loginHandler).Methods("POST")
}

func (controller *HttpController) loginHandler(w http.ResponseWriter, r *http.Request) {
	var data dto.LoginRequest
	if err := controller.Bind(r.Body, &data); err != nil {
		controller.sendErrorBadRequestResponse(w, err)
		return
	}

	res, err := controller.services.AuthService.Login(&data)
	if err != nil {
		controller.sendErrorResponse(w, err)
		return
	}

	controller.sendSuccessResponse(w, res)
}
