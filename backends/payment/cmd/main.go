package main

import (
	di "github.com.br/gibranct/payment/internal/infra/DI"
	"github.com.br/gibranct/payment/internal/infra/http"
)

func main() {
	httpService := http.NewHttpServer(di.NewProcessPayment())

	httpService.StartServer()
}
