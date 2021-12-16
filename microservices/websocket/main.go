package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"websocket/app/controllers"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	log.Println("Websocket microservice starting...")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	controllersConfig := &controllers.Config{
		HttpListerningAddres: os.Getenv("SERVICE_HTTP_ADDR"),
		HttpListerningPort:   os.Getenv("SERVICE_HTTP_PORT"),
		NatsServerAdress:     os.Getenv("NATS_ADDR"),
	}
	controllers, err := controllers.Create(controllersConfig)
	if err != nil {
		panic(err)
	}
	controllers.Run()
	defer controllers.Stop()

	log.Println("Websocket microservice started.")

	go syscallWait(cancelFunc)
	<-ctx.Done()
	log.Println("Websocket microservice stoping...")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}
