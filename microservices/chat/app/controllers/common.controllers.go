package controllers

import (
	nats_controller "chat/app/controllers/nats-controller"
	"chat/app/services"
)

type Controllers struct {
	NatsController *nats_controller.NatsController
}

type Config struct {
	NatsServerAdress string
	ErrorChannel     chan error
}

func Create(services *services.Services, conf *Config) (*Controllers, error) {
	natsControllerConfig := &nats_controller.Config{
		NatsServerAdress: conf.NatsServerAdress,
		ErrorChannel:     conf.ErrorChannel,
	}
	natsController, err := nats_controller.Create(services, natsControllerConfig)
	if err != nil {
		return nil, err
	}

	controllers := &Controllers{
		NatsController: natsController,
	}

	return controllers, nil
}

func (controllers *Controllers) Run() error {
	if err := controllers.NatsController.Run(); err != nil {
		return err
	}

	return nil
}

func (controllers *Controllers) Stop() {
	controllers.NatsController.Stop()
}
