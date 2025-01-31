package http

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com.br/gibranct/account/internal/application"
	"github.com.br/gibranct/account/internal/infra/controller"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	app     *application.Application
	handler *echo.Echo
}

func (http *HttpServer) StartServer() error {
	return http.handler.Start(os.Getenv("HOST"))
}

func (http *HttpServer) SetUpRoutes() {
	e := echo.New()

	http.handler = e

	accountCtrl := controller.NewAccountController(http.app.AccountService)

	e.Use(echoprometheus.NewMiddleware("account_api"))

	e.GET("/metrics", echoprometheus.NewHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1M"))

	e.POST("/v1/sign-up", accountCtrl.SignUpHandler)
	e.GET("/v1/accounts/:id", accountCtrl.GetAccountByIDHandler)
}

func (http *HttpServer) StopServer() {
	if err := http.handler.Server.Shutdown(context.Background()); err != nil {
		log.Fatalln(err)
	}
}

func (http *HttpServer) GetHandler() http.Handler {
	return http.handler
}

func NewHttpServer(app *application.Application) *HttpServer {
	return &HttpServer{
		app: app,
	}
}
