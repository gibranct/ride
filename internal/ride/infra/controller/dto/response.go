package dto

import (
	"time"

	"github.com.br/gibranct/ride/internal/ride/application/usecase"
)

type RideResponseDto struct {
	RideId      string  `json:"ride_id"`
	PassengerId string  `json:"passenger_id"`
	DriverId    string  `json:"driver_id"`
	Status      string  `json:"status"`
	Fare        float64 `json:"fare"`
	Distance    float64 `json:"distance"`
	FromLat     float64 `json:"from_lat"`
	Date        string  `json:"date"`
}

func FromRideToRideResponseDto(output usecase.GetRideOutput) RideResponseDto {
	return RideResponseDto{
		RideId:      output.RideId,
		PassengerId: output.PassengerId,
		DriverId:    output.DriverId,
		Status:      output.Status,
		Fare:        output.Fare,
		Distance:    output.Distance,
		FromLat:     output.FromLat,
		Date:        output.Date.Format(time.RFC3339),
	}
}
