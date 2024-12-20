package main

import (
	"github.com.br/gibranct/ride/internal/ride/application"
	"github.com.br/gibranct/ride/internal/ride/infra/http"
)

func main() {
	httpService := http.NewHttpServer(application.NewApplication())

	httpService.StartServer()
}
