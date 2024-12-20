package usecase

import (
	"fmt"

	"github.com.br/gibranct/ride/internal/ride/infra/repository"
)

type StartRideInput struct {
	RideId string
}

type StartRide struct {
	rideRepository repository.RideRepository
}

func (ar *StartRide) Execute(input StartRideInput) error {
	ride, err := ar.rideRepository.GetRideByID(input.RideId)
	if err != nil {
		return fmt.Errorf("ride not found: %s", err)
	}
	if err = ride.Start(); err != nil {
		return err
	}
	return ar.rideRepository.UpdateRide(*ride)
}

func NewStartRideUseCase(rideRepo repository.RideRepository) *StartRide {
	return &StartRide{
		rideRepository: rideRepo,
	}
}
