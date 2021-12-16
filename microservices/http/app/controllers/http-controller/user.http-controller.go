package http_controller

import (
	"http/app/api/dto"
	"net/http"
)

func (controller *HttpController) addUserHandlers() {
	subRouter := controller.router.PathPrefix("/users").Subrouter()

	subRouter.HandleFunc("/create", controller.handlerCreateUser).Methods("POST")
}

func (controller *HttpController) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var data dto.CreateUserRequest
	if err := controller.Bind(r.Body, &data); err != nil {
		controller.sendErrorBadRequestResponse(w, err)
		return
	}

	if err := controller.ValidateOrSendError(w, data); err != nil {
		return
	}

	res, err := controller.services.UserService.CreateUser(&data)
	if err != nil {
		controller.sendErrorResponse(w, err)
		return
	}

	controller.sendSuccessResponse(w, res)
}
