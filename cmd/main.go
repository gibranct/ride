package main

import (
	"github.com.br/gibranct/ride/cmd/application"
	"github.com.br/gibranct/ride/cmd/infra/http"
)

func main() {
	httpService := http.NewHttpServer(application.NewApplication())

	httpService.StartServer()
}
