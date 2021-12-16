package http_controller

import (
	"context"
	"http/app/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpController struct {
	conf     *Config
	server   *http.Server
	router   *mux.Router
	services *services.Services
}

type Config struct {
	ListeningAddr string
	ListeningPort string
}

func Create(conf *Config, services *services.Services) (*HttpController, error) {
	router := mux.NewRouter()

	controller := &HttpController{
		conf:     conf,
		router:   router,
		services: services,
	}

	controller.addAuthHandlers()
	controller.addUserHandlers()

	return controller, nil
}

func (controller *HttpController) Run() error {
	controller.server = &http.Server{
		Addr:    controller.conf.ListeningAddr + ":" + controller.conf.ListeningPort,
		Handler: controller.router,
	}

	if err := controller.server.ListenAndServe(); err != nil {
		log.Println("-> ERROR on run http server:", err)
		return err
	}

	return nil
}

func (controller *HttpController) Stop() error {
	return controller.server.Shutdown(context.Background())
}
