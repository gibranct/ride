package main

import (
	"log"

	"github.com.br/gibranct/account/internal/application"
	"github.com.br/gibranct/account/internal/infra/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	httpService := http.NewHttpServer(application.NewApplication())

	httpService.SetUpRoutes()
	err = httpService.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
