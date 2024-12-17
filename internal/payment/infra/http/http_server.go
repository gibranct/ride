package http

import (
	"github.com.br/gibranct/ride/internal/payment/application"
	"github.com.br/gibranct/ride/internal/payment/infra/controller"
	"github.com/labstack/echo/v4"
)

const port = "127.0.0.1:3333"

type HttpServer struct {
	app *application.Application
}

func (http *HttpServer) StartServer() {
	e := echo.New()

	paymentCtrl := controller.NewPaymentController(http.app.PaymentService)

	e.POST("/process_payment", paymentCtrl.ProcessPaymentHandler)

	e.Logger.Fatal(e.Start(port))
}

func NewHttpServer(app *application.Application) *HttpServer {
	return &HttpServer{
		app: app,
	}
}
