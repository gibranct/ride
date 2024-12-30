package main

import (
	"github.com.br/gibranct/ride/internal/application"
	"github.com.br/gibranct/ride/internal/infra/controller"
	"github.com.br/gibranct/ride/internal/infra/http"
)

func main() {
	httpService := http.NewHttpServer(controller.NewRideController(
		application.NewApplication().RideService,
	))

	httpService.StartServer()
}
