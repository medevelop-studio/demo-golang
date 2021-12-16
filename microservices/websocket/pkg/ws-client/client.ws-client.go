package ws_client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type ClientStatusCode byte

const (
	CLIENT_STATUS_NEW ClientStatusCode = iota + 1
	CLIENT_STATUS_PENDING
	CLIENT_STATUS_READY
	CLIENT_STATUS_STOP
)

const (
	DEFAULT_WRITE_WAIT       = 60 * time.Second
	DEFAULT_PONG_WAIT        = 60 * time.Second
	DEFAULT_PING_PERIOD      = (DEFAULT_PONG_WAIT * 9) / 10
	DEFAULT_MAX_MESSAGE_SIZE = 512
)

type SendFunction func(data []byte) error
type HandlerFunction func(data []byte, connId string, send func(data []byte) error) error

type Client struct {
	conn *websocket.Conn

	ConnId               string
	Status               ClientStatusCode
	sendToClient         chan []byte
	sendToClientBuffered chan []byte

	handlerMessageCallback   HandlerFunction
	handlerCloseConnCallback HandlerFunction
	handlerErrorCallback     HandlerFunction

	ctx        context.Context
	cancelFunc context.CancelFunc
}

func CreateClient(connId string, conn *websocket.Conn) *Client {
	return &Client{
		ConnId:               connId,
		conn:                 conn,
		Status:               CLIENT_STATUS_NEW,
		sendToClient:         make(chan []byte),
		sendToClientBuffered: make(chan []byte, 5),
	}
}

func (client *Client) SetHandlerMessage(
	handler func(data []byte, connId string, send func(data []byte) error) error,
) {
	client.handlerMessageCallback = handler
}
func (client *Client) SetHandlerCloseConn(handler HandlerFunction) {
	client.handlerCloseConnCallback = handler
}
func (client *Client) SetHandlerError(handler HandlerFunction) {
	client.handlerErrorCallback = handler
}

func (client *Client) Run() error {
	if client.handlerMessageCallback == nil {
		return ErrClientHandlerMessageNotSet
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	client.ctx = ctx
	client.cancelFunc = cancelFunc

	client.conn.SetCloseHandler(func(_ int, _ string) error {
		client.abort()
		if client.handlerCloseConnCallback != nil {
			client.handlerCloseConnCallback(client.createReqCtx(nil))
		}
		return nil
	})

	client.conn.SetReadLimit(DEFAULT_MAX_MESSAGE_SIZE)
	if err := client.conn.SetReadDeadline(time.Now().Add(DEFAULT_PONG_WAIT)); err != nil {
		return err
	}
	client.conn.SetPongHandler(func(string) error {
		return client.conn.SetReadDeadline(time.Now().Add(DEFAULT_PONG_WAIT))
	})

	go client.reader()
	go client.writer()

	return nil
}
func (client *Client) Stop() error {
	client.Status = CLIENT_STATUS_STOP
	client.conn.WriteMessage(websocket.CloseMessage, []byte{})
	client.cancelFunc()

	return client.conn.Close()
}

func (client *Client) SendByteMessage(data []byte) (err error) {
	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			client.handlerPanic(recoveryMessage)
			err = ErrClientCloseConnect
		}
	}()
	if client.Status == CLIENT_STATUS_STOP {
		return ErrClientCloseConnect
	}

	client.sendToClient <- data
	return
}

func (client *Client) SendMessage(data interface{}) (err error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return client.SendByteMessage(byteData)
}

func (client *Client) GetClientContext() ([]byte, string, SendFunction) {
	return client.createReqCtx(nil)
}

func (client *Client) createReqCtx(data []byte) ([]byte, string, SendFunction) {
	return data, client.ConnId, client.SendByteMessage
}

func (client *Client) abort() {
	client.Status = CLIENT_STATUS_STOP
	client.cancelFunc()
	client.conn.Close()
}

func (client *Client) handlerMessage(message []byte) {
	if client.handlerMessageCallback != nil {
		client.handlerMessageCallback(client.createReqCtx(message))
	}
}

func (client *Client) handlerError(err error, isInitiatedInRider bool) {
	if client.Status == CLIENT_STATUS_STOP {
		return
	}

	if err == websocket.ErrReadLimit {
		if client.handlerErrorCallback != nil {
			client.handlerErrorCallback(client.createReqCtx([]byte(err.Error())))
		}

		return
	}

	client.abort()
	if client.handlerCloseConnCallback != nil {
		client.handlerCloseConnCallback(client.createReqCtx([]byte(err.Error())))
	}
}

func (client *Client) handlerPanic(panic interface{}) {
	log.Println("[handlerPanic] - PANIC: ", panic)

	client.abort()
	if client.handlerErrorCallback != nil {
		client.handlerErrorCallback(client.createReqCtx([]byte(fmt.Sprint(panic))))
	}
}
