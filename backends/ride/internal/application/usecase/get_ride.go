package usecase

import (
	"time"

	domain "github.com.br/gibranct/ride/internal/domain/entity"
	"github.com.br/gibranct/ride/internal/domain/service"
	"github.com.br/gibranct/ride/internal/infra/repository"
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
	Fare        float64
	Date        *time.Time
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
	var distance float64
	if ride.GetStatus() == domain.COMPLETED_RIDE_STATUS {
		distance = ride.GetDistance()
	} else {
		var newPositions []service.Position
		for _, pos := range positions {
			newPositions = append(newPositions, pos)
		}
		distance = service.NewDistanceCalculator().CalculateByPositions(newPositions)
	}

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
		Fare:        ride.GetFare(),
		Date:        ride.GetDate(),
	}, nil
}

func NewGetRideUseCase(rideRepo repository.RideRepository, positionRepository repository.PositionRepository) *GetRide {
	return &GetRide{
		rideRepository:     rideRepo,
		positionRepository: positionRepository,
	}
}
