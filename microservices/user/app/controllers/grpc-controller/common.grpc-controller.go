package grpc_controller

import (
	"log"
	"net"
	"user/app/services"
	"user/proto"

	"google.golang.org/grpc"
)

type GrpcController struct {
	server      *grpc.Server
	netListener net.Listener
}

type Config struct {
	Protocol string
	Port     string
}

func Create(services *services.Services, conf *Config) (*GrpcController, error) {
	listener, err := createNetListener(conf.Protocol, conf.Port)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()

	controller := &GrpcController{
		netListener: listener,
		server:      server,
	}

	controller.registerServices(services)

	return controller, nil
}

func (controller *GrpcController) Run() error {
	if err := controller.server.Serve(controller.netListener); err != nil {
		log.Println("-> ERROR on run grpc server:", err)
		return err
	}

	return nil
}

func createNetListener(addr, port string) (net.Listener, error) {
	listener, err := net.Listen(addr, ":"+port)
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func (controller *GrpcController) registerServices(services *services.Services) {
	proto.RegisterAuthServiceServer(controller.server, &AuthServiceController{services: services})
	proto.RegisterUserServiceServer(controller.server, &UserServiceController{services: services})
}
