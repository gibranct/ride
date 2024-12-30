package main

import (
	"github.com.br/gibranct/account/internal/application"
	"github.com.br/gibranct/account/internal/infra/http"
)

func main() {
	httpService := http.NewHttpServer(application.NewApplication())

	httpService.StartServer()
}
