package controllers

import (
	http_controller "http/app/controllers/http-controller"
	"http/app/services"
)

type Controllers struct {
	HttpController *http_controller.HttpController
}

type Config struct {
	HttpListerningAddres string
	HttpListerningPort   string
}

func Create(conf *Config, services *services.Services) (*Controllers, error) {
	httpControllerConfig := &http_controller.Config{
		ListeningAddr: conf.HttpListerningAddres,
		ListeningPort: conf.HttpListerningPort,
	}
	httpController, err := http_controller.Create(httpControllerConfig, services)
	if err != nil {
		return nil, err
	}

	controllers := &Controllers{
		HttpController: httpController,
	}

	return controllers, nil
}

func (controllers *Controllers) Run() error {
	go controllers.HttpController.Run()

	return nil
}

func (controllers *Controllers) Stop() error {
	return controllers.HttpController.Stop()
}
