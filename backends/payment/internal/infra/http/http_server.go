package http

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com.br/gibranct/payment/internal/infra/controller"
	"github.com.br/gibranct/payment/internal/infra/queue"
	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	processPayment usecase.IProcessPayment
	handler        *echo.Echo
}

func (http *HttpServer) StartServer() {
	http.handler.Logger.Fatal(http.handler.Start(os.Getenv("HOST")))
}

func (http *HttpServer) SetUpRoutes() {
	e := echo.New()

	http.handler = e

	paymentCtrl := controller.NewPaymentController(http.processPayment)
	controller.NewQueueController(http.processPayment, queue.NewRabbitMQAdapter())

	e.POST("/process_payment", paymentCtrl.ProcessPaymentHandler)
}

func (http *HttpServer) GetHandler() http.Handler {
	return http.handler
}

func (http *HttpServer) StopServer() {
	if err := http.handler.Server.Shutdown(context.Background()); err != nil {
		log.Fatalln(err)
	}
}

func NewHttpServer(processPayment usecase.IProcessPayment) *HttpServer {
	return &HttpServer{
		processPayment: processPayment,
	}
}
