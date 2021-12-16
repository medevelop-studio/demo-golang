package nats_controller

import (
	"errors"
	"log"
	"time"
	"user/app/services"

	"github.com/nats-io/nats.go"
)

type NatsController struct {
	conf     *Config
	natsConn *nats.Conn
	services *services.Services
}

type Config struct {
	NatsServerAdress string
}

func Create(services *services.Services, conf *Config) (*NatsController, error) {
	controller := &NatsController{
		conf:     conf,
		services: services,
	}

	return controller, nil
}

func (controller *NatsController) Run() error {
	natsConn, err := controller.connectToNats()
	if err != nil {
		return err
	}

	controller.natsConn = natsConn

	controller.createChatController()

	log.Println("-> NATS Connected.")
	return nil
}

func (controller *NatsController) Stop() {
	controller.natsConn.Close()
	log.Println("-> NATS Closed connection.")
}

func (controller *NatsController) connectToNats() (*nats.Conn, error) {
	var conn *nats.Conn
	var err error

	for i := 0; i < 3; i++ {
		conn, err = nats.Connect(controller.conf.NatsServerAdress)
		if err == nil {
			break
		}

		log.Println("-> ERROR - Connect to NATS. Waiting for reconnection...")
		time.Sleep(time.Duration(5) * time.Second)
	}

	if conn == nil {
		return nil, errors.New("-> ERROR - Connect to NATS")
	}

	return conn, nil
}
