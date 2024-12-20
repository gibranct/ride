package main

import (
	"github.com.br/gibranct/ride/internal/account/application"
	"github.com.br/gibranct/ride/internal/account/infra/http"
)

func main() {
	httpService := http.NewHttpServer(application.NewApplication())

	httpService.StartServer()
}
