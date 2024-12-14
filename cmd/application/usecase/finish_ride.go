package usecase

import (
	"fmt"

	"github.com.br/gibranct/ride/cmd/infra/repository"
)

type FinishRideInput struct {
	RideId string
}

type FinishRide struct {
	rideRepository     repository.RideRepository
	positionRepository repository.PositionRepository
}

func (ar *FinishRide) Execute(input FinishRideInput) error {
	ride, err := ar.rideRepository.GetRideByID(input.RideId)
	if err != nil {
		return fmt.Errorf("ride not found: %s", err)
	}
	positions, err := ar.positionRepository.GetPositionsByRideId(input.RideId)
	if err != nil {
		return fmt.Errorf("positions not found: %s", err)
	}
	err = ride.Finish(positions)
	if err != nil {
		return err
	}
	return ar.rideRepository.UpdateRide(*ride)
}

func NewFinishRideUseCase(
	rideRepo repository.RideRepository,
	positionRepo repository.PositionRepository,
) *FinishRide {
	return &FinishRide{
		rideRepository:     rideRepo,
		positionRepository: positionRepo,
	}
}
