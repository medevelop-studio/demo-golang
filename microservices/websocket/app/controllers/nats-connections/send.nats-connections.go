package nats_connections

import (
	"encoding/json"
	"websocket/app/api/subjects"
)

func (conns *NatsConnections) SendLoginRequest(req interface{}) error {
	return conns.publish(subjects.LOGIN_REQUEST, req)
}

func (conns *NatsConnections) SendMessageRequest(req interface{}) error {
	return conns.publish(subjects.CHAT_REQUEST, req)
}

func (conns *NatsConnections) publish(sub string, msg interface{}) error {
	byteData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return conns.natsConn.Publish(sub, byteData)
}
