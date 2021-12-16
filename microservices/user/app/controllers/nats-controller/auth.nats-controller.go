package nats_controller

import (
	"user/app/api/codes"
	"user/app/api/dto"
	"user/app/api/subjects"
	"user/pkg/router"
)

func (controller *NatsController) createChatController() {
	router := router.CreateRouter()

	router.SetParser(controller.parseMessageCode)
	router.SetErrorHandler(controller.handlerError)

	router.AddHandler(codes.LoginRequest, controller.handlerLoginRequest)

	controller.addSubscriber(subjects.LOGIN_REQUEST, subjects.LOGIN_QUEUE, router.NewMessage)
}

func (controller *NatsController) handlerLoginRequest(data []byte, subjectId string) error {
	var req dto.LoginWTNRequest
	if err := controller.bind(data, &req); err != nil {
		return err
	}

	loginData, err := controller.services.AuthService.LoginUser(req.Login, req.Password)
	if err != nil {
		res := dto.LoginNTWResponse{
			ResponseGeneral: dto.CreateResponse(codes.LoginResponse),
			ConnectionId:    req.ConnectionId,
		}
		res.Error = dto.ConvertErrToError(err)
		res.Request = req
		return controller.send(subjects.WS_RESPONSE, res)
	}

	res := &dto.LoginNTWResponse{
		ResponseGeneral: dto.CreateResponse(codes.LoginResponse),
		ConnectionId:    req.ConnectionId,
		Data:            loginData,
	}

	return controller.send(subjects.WS_RESPONSE, res)
}
