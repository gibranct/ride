package controller

import (
	"github.com.br/gibranct/ride/internal/ride/application"
	"github.com/labstack/echo/v4"
)

type RideController struct {
	rideService application.RideService
}

func (rideCtrl *RideController) SignUpHandler(c echo.Context) error {
	return nil
}

func (rideCtrl *RideController) GetAccountByIDHandler(c echo.Context) error {
	return nil
}

func NewRideController(
	rideService *application.RideService,
) *RideController {
	return &RideController{
		rideService: *rideService,
	}
}
