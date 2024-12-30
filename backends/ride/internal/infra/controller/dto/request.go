package dto

import (
	"github.com.br/gibranct/ride/internal/application/usecase"
)

type RideRequestDto struct {
	PassengerId string  `json:"passenger_id"`
	FromLat     float64 `json:"from_lat"`
	FromLong    float64 `json:"from_long"`
	ToLat       float64 `json:"to_lat"`
	ToLong      float64 `json:"to_long"`
}

func (dto *RideRequestDto) ToRequestRideInput() usecase.RequestRideInput {
	return usecase.RequestRideInput{
		PassengerId: dto.PassengerId,
		FromLat:     dto.FromLat,
		FromLong:    dto.FromLong,
		ToLat:       dto.ToLat,
		ToLong:      dto.ToLong,
	}
}

type RideStartRequestDto struct {
	RideId string `json:"ride_id"`
}

func (dto *RideStartRequestDto) ToStartRideInput() usecase.StartRideInput {
	return usecase.StartRideInput{
		RideId: dto.RideId,
	}
}

type RideAcceptRequestDto struct {
	RideId   string `json:"ride_id"`
	DriverId string `json:"driver_id"`
}

func (dto *RideAcceptRequestDto) ToAcceptRideInput() usecase.AcceptRideInput {
	return usecase.AcceptRideInput{
		RideId:   dto.RideId,
		DriverId: dto.DriverId,
	}
}

type RideFinishRequestDto struct {
	RideId string `json:"ride_id"`
}

func (dto *RideFinishRequestDto) ToFinishRideInput() usecase.FinishRideInput {
	return usecase.FinishRideInput{
		RideId: dto.RideId,
	}
}
