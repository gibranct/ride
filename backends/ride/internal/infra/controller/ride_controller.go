package controller

import (
	"net/http"

	"github.com.br/gibranct/ride/internal/application"
	"github.com.br/gibranct/ride/internal/infra/controller/dto"
	"github.com/labstack/echo/v4"
)

type RideController struct {
	rideService application.RideService
}

func (rideCtrl *RideController) RequestRide(c echo.Context) error {
	var input dto.RideRequestDto

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	output, err := rideCtrl.rideService.RequestRide.Execute(input.ToRequestRideInput())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"ride_id": output.RideId})
}

func (rideCtrl *RideController) StartRide(c echo.Context) error {
	var input dto.RideStartRequestDto

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err := rideCtrl.rideService.StartRide.Execute(input.ToStartRideInput())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (rideCtrl *RideController) AcceptRide(c echo.Context) error {
	var input dto.RideAcceptRequestDto

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err := rideCtrl.rideService.AcceptRide.Execute(input.ToAcceptRideInput())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (rideCtrl *RideController) FinishRide(c echo.Context) error {
	var input dto.RideFinishRequestDto

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err := rideCtrl.rideService.FinishRide.Execute(input.ToFinishRideInput())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (rideCtrl *RideController) GetRide(c echo.Context) error {
	rideId := c.Param("id")
	output, err := rideCtrl.rideService.GetRide.Execute(rideId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.FromRideToRideResponseDto(*output))
}

func NewRideController(
	rideService *application.RideService,
) *RideController {
	return &RideController{
		rideService: *rideService,
	}
}
