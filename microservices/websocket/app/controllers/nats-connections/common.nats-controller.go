package nats_connections

import (
	"errors"
	"log"
	"time"
	"websocket/pkg/router"

	"github.com/nats-io/nats.go"
)

type NatsConnections struct {
	conf          *Config
	natsConn      *nats.Conn
	serverRouter  *router.Router
	messageRouter *router.Router
}

type Config struct {
	NatsServerAdress string
}

func Create(conf *Config) (*NatsConnections, error) {
	controller := &NatsConnections{
		conf: conf,
	}

	return controller, nil
}

func (conns *NatsConnections) SetServerRouter(router *router.Router) {
	conns.serverRouter = router
}

func (conns *NatsConnections) SetMessageRouter(router *router.Router) {
	conns.messageRouter = router
}

func (conns *NatsConnections) Run() error {
	natsConn, err := conns.connectToNats()
	if err != nil {
		return err
	}

	conns.natsConn = natsConn

	if err := conns.addSubscribers(); err != nil {
		natsConn.Close()
		return err
	}

	log.Println("-> NATS Connected.")
	return nil
}

func (conns *NatsConnections) Stop() {
	conns.natsConn.Close()
	log.Println("-> NATS Closed connection.")
}

func (conns *NatsConnections) connectToNats() (*nats.Conn, error) {
	var conn *nats.Conn
	var err error

	for i := 0; i < 3; i++ {
		conn, err = nats.Connect(conns.conf.NatsServerAdress)
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
