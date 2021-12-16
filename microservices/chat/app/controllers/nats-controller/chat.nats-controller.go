package nats_controller

import (
	"chat/app/api/codes"
	"chat/app/api/dto"
	"chat/app/api/subjects"

	"github.com/nats-io/nats.go"
)

func (controller *NatsController) createChatController() {
	router := CreateRouter()

	router.SetParser(controller.parseMessageCode)
	router.SetErrorHandler(controller.handlerError)

	router.AddHandler(codes.ChatMessageRequest, controller.handlerChatMessage)

	controller.addSubscriber(
		subjects.CHAT_REQUEST,
		subjects.CHAT_REQUEST_QUERY,
		router.NewMessage,
	)
}

func (controller *NatsController) handlerChatMessage(message *nats.Msg, subjectId string) error {
	var req dto.ChatMessageRequest
	if err := controller.bind(message.Data, &req); err != nil {
		return err
	}

	data, err := controller.services.MessageService.NewMessageToChat(req.Data)
	if err != nil {
		return controller.sendErrorToUser(req.Data.UserId, err, req)
	}

	res := &dto.ChatMessageResponse{
		ResponseGeneral: dto.CreateResponse(codes.ChatMessageResponse),
		Data:            data,
	}

	return controller.send(subjects.MESSAGE_RESPONSE, res)
}
