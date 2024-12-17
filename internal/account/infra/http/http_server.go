package http

import (
	"github.com.br/gibranct/ride/internal/account/application"
	"github.com.br/gibranct/ride/internal/account/infra/controller"
	"github.com/labstack/echo/v4"
)

const port = "127.0.0.1:3333"

type HttpServer struct {
	app *application.Application
}

func (http *HttpServer) StartServer() {
	e := echo.New()

	accountCtrl := controller.NewAccountController(http.app.AccountService)

	e.POST("/sign-up", accountCtrl.SignUpHandler)
	e.GET("/v1/accounts/:id", accountCtrl.GetAccountByIDHandler)

	e.Logger.Fatal(e.Start(port))
}

func NewHttpServer(app *application.Application) *HttpServer {
	return &HttpServer{
		app: app,
	}
}
