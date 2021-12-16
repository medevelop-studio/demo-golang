package http_controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type HttpController struct {
	conf           *Config
	server         *http.Server
	newConnHandler HandlerNewConnection
}

type Config struct {
	ListeningAddr string
	ListeningPort string
}

type HandlerNewConnection func(*websocket.Conn) error

func Create(conf *Config) (*HttpController, error) {
	controller := &HttpController{
		conf: conf,
	}

	controller.addNewConnectionHandler()

	return controller, nil
}

func (controller *HttpController) Run() error {
	controller.server = &http.Server{
		Addr: controller.conf.ListeningAddr + ":" + controller.conf.ListeningPort,
	}

	if err := controller.server.ListenAndServe(); err != nil {
		log.Println("-> ERROR on run http server:", err)
		return err
	}

	return nil
}

func (controller *HttpController) SetNewConnectionHandler(handler HandlerNewConnection) {
	controller.newConnHandler = handler
}

func (controller *HttpController) Stop() error {
	return controller.server.Shutdown(context.Background())
}
