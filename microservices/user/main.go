package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"user/app/controllers"
	"user/app/services"
	"user/app/store"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	log.Println("User microservice starting...")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	repository, err := store.CreateMockStore()
	if err != nil {
		panic(err)
	}
	defer repository.Close()

	authTokenLifeTime, err := strconv.ParseInt(os.Getenv("TOKEN_LIFE_TIME"), 10, 64)
	if err != nil {
		panic(err)
	}
	servicesConfig := &services.Config{
		AuthTokenSecret:   os.Getenv("TOKEN_SECRET"),
		AuthTokenLifeTime: authTokenLifeTime,
	}
	services, err := services.Create(repository, servicesConfig)
	if err != nil {
		panic(err)
	}

	controllersConfig := &controllers.Config{
		GrpcServerListerProtocol: os.Getenv("SERVICE_GRPC_PROTOCOL"),
		GrpcServerListerPort:     os.Getenv("SERVICE_GRPC_PORT"),
		NatsServerAdress:         os.Getenv("NATS_ADDR"),
	}
	controllers, err := controllers.Create(services, controllersConfig)
	if err != nil {
		panic(err)
	}
	controllers.Run()
	defer controllers.Stop()

	log.Println("User microservice started.")

	go syscallWait(cancelFunc)
	<-ctx.Done()
	log.Println("User microservice stoping...")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}
