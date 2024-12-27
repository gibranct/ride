package http

import (
	"github.com.br/gibranct/ride/internal/ride/infra/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const port = "127.0.0.1:3000"

type HttpServer struct {
	rideController *controller.RideController
}

func (http *HttpServer) StartServer() {
	e := echo.New()

	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.Logger())

	e.POST("/v1/ride/request", http.rideController.RequestRide)
	e.POST("/v1/ride/accept", http.rideController.AcceptRide)
	e.POST("/v1/ride/start", http.rideController.StartRide)
	e.POST("/v1/ride/finish", http.rideController.FinishRide)
	e.GET("/v1/ride/:id", http.rideController.GetRide)

	e.Logger.Fatal(e.Start(port))
}

func NewHttpServer(rideController *controller.RideController) *HttpServer {
	return &HttpServer{
		rideController: rideController,
	}
}
