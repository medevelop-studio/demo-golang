package nats_controller

import (
	"chat/app/api/codes"
	"chat/app/api/dto"
	"chat/app/api/subjects"
	"encoding/json"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
)

type HandlerFunc func(*nats.Msg, string) error

func (controller *NatsController) addSubscriber(
	sub string,
	queue string,
	handler HandlerFunc,
) error {
	_, err := controller.natsConn.QueueSubscribe(sub, queue, func(message *nats.Msg) {
		defer func() {
			if msg := recover(); msg != nil {
				log.Println("[AddSubscriber] - Call recover - ERROR: ", msg)
			}
		}()

		var subId string
		position := strings.Index(message.Sub.Subject, "-")
		if position > 0 {
			subId = message.Sub.Subject[position+1:]
		}

		handler(message, subId)
	})

	if err != nil {
		return err
	}

	return nil
}

func (controller *NatsController) parseMessageCode(data []byte) (codes.MessageTypeCode, error) {
	var message dto.RequestGeneral

	if err := controller.bind(data, &message); err != nil {
		return 0, err
	}

	return message.Type, nil
}

func (controller *NatsController) bind(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}

func (controller *NatsController) handlerError(message *nats.Msg, subjectID string, err error) {
	log.Println("-> ERROR route: ", err)
}

func (controller *NatsController) sendErrorToUser(userId string, err error, req interface{}) error {
	subject := subjects.GetUserResponseSub(userId)

	res := &dto.ResponseGeneral{
		Error:   dto.ConvertErrToError(err),
		Request: req,
	}

	return controller.send(subject, res)
}

func (controller *NatsController) send(subject string, data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return controller.natsConn.Publish(subject, byteData)
}
