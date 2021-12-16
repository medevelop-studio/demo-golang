package nats_connections

import (
	"log"
	"strings"

	"github.com/nats-io/nats.go"
)

type HandlerFunc func([]byte, string, func(data []byte) error) error

func (conns *NatsConnections) addSubscriber(sub string, handler HandlerFunc) error {
	_, err := conns.natsConn.Subscribe(sub, func(message *nats.Msg) {
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

		handler(message.Data, subId, nil)
	})

	if err != nil {
		return err
	}

	return nil
}
