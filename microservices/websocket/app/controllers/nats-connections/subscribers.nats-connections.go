package nats_connections

import "websocket/app/api/subjects"

func (conns *NatsConnections) addSubscribers() error {
	if err := conns.addSubscriber(subjects.WS_RESPONSE, conns.serverRouter.NewMessage); err != nil {
		return err
	}

	if err := conns.addSubscriber(subjects.MESSAGE_RESPONSE, conns.messageRouter.NewMessage); err != nil {
		return err
	}

	return nil
}
