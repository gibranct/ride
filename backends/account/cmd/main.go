package main

import (
	"log"

	"github.com.br/gibranct/account/internal/application"
	"github.com.br/gibranct/account/internal/infra/http"
)

func main() {
	httpService := http.NewHttpServer(application.NewApplication())

	log.Fatalln(httpService.StartServer())
}
