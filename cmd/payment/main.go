package main

import (
	di "github.com.br/gibranct/ride/internal/payment/infra/DI"
	"github.com.br/gibranct/ride/internal/payment/infra/http"
)

func main() {
	httpService := http.NewHttpServer(di.NewProcessPayment())

	httpService.StartServer()
}
