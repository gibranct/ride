package http

import (
	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com.br/gibranct/payment/internal/infra/controller"
	"github.com.br/gibranct/payment/internal/infra/queue"
	"github.com/labstack/echo/v4"
)

const port = "127.0.0.1:3002"

type HttpServer struct {
	processPayment usecase.IProcessPayment
}

func (http *HttpServer) StartServer() {
	e := echo.New()

	paymentCtrl := controller.NewPaymentController(http.processPayment)
	controller.NewQueueController(http.processPayment, queue.NewRabbitMQAdapter())

	e.POST("/process_payment", paymentCtrl.ProcessPaymentHandler)

	e.Logger.Fatal(e.Start(port))
}

func NewHttpServer(processPayment usecase.IProcessPayment) *HttpServer {
	return &HttpServer{
		processPayment: processPayment,
	}
}
