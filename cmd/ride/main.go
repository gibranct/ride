package main

import (
	"github.com.br/gibranct/ride/internal/ride/application"
	"github.com.br/gibranct/ride/internal/ride/infra/controller"
	"github.com.br/gibranct/ride/internal/ride/infra/http"
)

func main() {
	httpService := http.NewHttpServer(controller.NewRideController(
		application.NewApplication().RideService,
	))

	httpService.StartServer()
}
