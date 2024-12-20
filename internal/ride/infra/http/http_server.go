package http

import (
	"github.com.br/gibranct/ride/internal/ride/application"
	"github.com.br/gibranct/ride/internal/ride/infra/controller"
	"github.com/labstack/echo/v4"
)

const port = "127.0.0.1:3000"

type HttpServer struct {
	app *application.Application
}

func (http *HttpServer) StartServer() {
	e := echo.New()

	rideCtrl := controller.NewRideController(http.app.RideService)

	e.POST("/sign-up", rideCtrl.SignUpHandler)
	e.GET("/v1/accounts/:id", rideCtrl.GetAccountByIDHandler)

	e.Logger.Fatal(e.Start(port))
}

func NewHttpServer(app *application.Application) *HttpServer {
	return &HttpServer{
		app: app,
	}
}
