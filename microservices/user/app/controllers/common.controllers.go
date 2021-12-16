package controllers

import (
	grpc_controller "user/app/controllers/grpc-controller"
	nats_controller "user/app/controllers/nats-controller"
	"user/app/services"
)

type Controllers struct {
	grpcController *grpc_controller.GrpcController
	natsController *nats_controller.NatsController
}

type Config struct {
	GrpcServerListerProtocol string
	GrpcServerListerPort     string
	NatsServerAdress         string
}

func Create(services *services.Services, conf *Config) (*Controllers, error) {
	grpcControllerConfig := &grpc_controller.Config{
		Protocol: conf.GrpcServerListerProtocol,
		Port:     conf.GrpcServerListerPort,
	}
	grpcController, err := grpc_controller.Create(services, grpcControllerConfig)
	if err != nil {
		return nil, err
	}

	natsControllerConfig := &nats_controller.Config{
		NatsServerAdress: conf.NatsServerAdress,
	}
	natsController, err := nats_controller.Create(services, natsControllerConfig)
	if err != nil {
		return nil, err
	}

	controllers := &Controllers{
		grpcController: grpcController,
		natsController: natsController,
	}

	return controllers, nil
}

func (controllers *Controllers) Run() error {
	go controllers.grpcController.Run()

	return controllers.natsController.Run()
}

func (controllers *Controllers) Stop() {
}
