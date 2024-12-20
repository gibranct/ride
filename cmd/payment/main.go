package main

import (
	"github.com.br/gibranct/ride/internal/payment/application"
	"github.com.br/gibranct/ride/internal/payment/infra/http"
)

func main() {
	httpService := http.NewHttpServer(application.NewApplication())

	httpService.StartServer()
}
