package main

import (
	"context"
	"http/app/controllers"
	"http/app/services"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	log.Println("Http microservice starting...")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	servicesConfig := &services.Config{
		GRPCUserMicroserviceAddress: os.Getenv("USER_SERVICE_ADDR"),
	}
	services, err := services.Create(servicesConfig)
	if err != nil {
		panic(err)
	}

	controllersConfig := &controllers.Config{
		HttpListerningAddres: os.Getenv("SERVICE_HTTP_ADDR"),
		HttpListerningPort:   os.Getenv("SERVICE_HTTP_PORT"),
	}
	controllers, err := controllers.Create(controllersConfig, services)
	if err != nil {
		panic(err)
	}
	controllers.Run()
	defer controllers.Stop()
	log.Println("Http microservice started.")

	go syscallWait(cancelFunc)
	<-ctx.Done()
	log.Println("Http microservice stoping...")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}
