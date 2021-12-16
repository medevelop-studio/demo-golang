package main

import (
	"chat/app/controllers"
	"chat/app/services"
	"chat/app/store"
	"context"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	errorChannel := make(chan error)
	log.Println("Chat microservice starting...")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	repository, err := store.CreateMockStore()
	if err != nil {
		panic(err)
	}
	defer repository.Close()

	servicesConfig := &services.Config{
		GRPCUserMicroserviceAddress: os.Getenv("USER_SERVICE_ADDR"),
	}
	services, err := services.Create(repository, servicesConfig)
	if err != nil {
		panic(err)
	}

	controllersConfig := &controllers.Config{
		NatsServerAdress: os.Getenv("NATS_ADDR"),
		ErrorChannel:     errorChannel,
	}
	controllers, err := controllers.Create(services, controllersConfig)
	if err != nil {
		panic(err)
	}
	if err := controllers.Run(); err != nil {
		panic(err)
	}
	defer controllers.Stop()

	log.Println("Chat microservice started.")

	go errorListener(errorChannel, ctx)
	go syscallWait(cancelFunc)
	<-ctx.Done()
	log.Println("Chat microservice stoping...")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}

func errorListener(errorChannel chan error, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-errorChannel:
			log.Println("ERROR ->", err)
			log.Println(string(debug.Stack()))
		}
	}
}
