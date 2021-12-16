package controllers

import (
	http_controller "websocket/app/controllers/http-controller"
	nats_connections "websocket/app/controllers/nats-connections"
	ws_connections "websocket/app/controllers/ws-connections"
)

type Controllers struct {
	httpController  *http_controller.HttpController
	natsConnections *nats_connections.NatsConnections
	wsConnections   *ws_connections.WsConnections
}

type Config struct {
	HttpListerningAddres string
	HttpListerningPort   string
	NatsServerAdress     string
}

func Create(conf *Config) (*Controllers, error) {
	httpControllerConfig := &http_controller.Config{
		ListeningAddr: conf.HttpListerningAddres,
		ListeningPort: conf.HttpListerningPort,
	}
	httpController, err := http_controller.Create(httpControllerConfig)
	if err != nil {
		return nil, err
	}

	natsConnectionsConfig := &nats_connections.Config{
		NatsServerAdress: conf.NatsServerAdress,
	}
	natsConnections, err := nats_connections.Create(natsConnectionsConfig)
	if err != nil {
		return nil, err
	}

	wsConnections, err := ws_connections.Create()
	if err != nil {
		return nil, err
	}

	controllers := &Controllers{
		httpController:  httpController,
		natsConnections: natsConnections,
		wsConnections:   wsConnections,
	}

	controllers.setHandlers()

	return controllers, nil
}

func (controllers *Controllers) Run() error {
	go controllers.httpController.Run()

	if err := controllers.natsConnections.Run(); err != nil {
		return err
	}

	return nil
}

func (controllers *Controllers) Stop() error {
	return controllers.httpController.Stop()
}
