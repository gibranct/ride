package main

import (
	"log"

	di "github.com.br/gibranct/payment/internal/infra/DI"
	"github.com.br/gibranct/payment/internal/infra/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpService := http.NewHttpServer(di.NewProcessPayment())

	httpService.SetUpRoutes()
	httpService.StartServer()
}
