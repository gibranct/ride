package usecase

import (
	"github.com.br/gibranct/ride/cmd/domain"
	"github.com.br/gibranct/ride/cmd/infra/repository"
)

type GetRide struct {
	rideRepository     repository.RideRepository
	positionRepository repository.PositionRepository
}

type GetRideOutput struct {
	RideId      string
	PassengerId string
	DriverId    string
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
	Status      string
	Positions   []domain.Position
	Distance    float64
}

func (gr *GetRide) Execute(rideId string) (*GetRideOutput, error) {
	ride, err := gr.rideRepository.GetRideByID(rideId)
	if err != nil {
		return nil, err
	}
	positions, err := gr.positionRepository.GetPositionsByRideId(rideId)
	if err != nil {
		return nil, err
	}
	distance := ride.GetDistance(positions)

	return &GetRideOutput{
		RideId:      ride.GetRideId(),
		PassengerId: ride.GetPassengerId(),
		DriverId:    ride.GetDriverId(),
		FromLat:     ride.GetFromCoord().GetLat(),
		FromLong:    ride.GetFromCoord().GetLong(),
		ToLat:       ride.GetToCoord().GetLat(),
		ToLong:      ride.GetToCoord().GetLong(),
		Status:      ride.GetStatus(),
		Positions:   positions,
		Distance:    distance,
	}, nil
}

func NewGetRideUseCase(rideRepo repository.RideRepository, positionRepository repository.PositionRepository) *GetRide {
	return &GetRide{
		rideRepository:     rideRepo,
		positionRepository: positionRepository,
	}
}
