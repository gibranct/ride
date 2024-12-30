package usecase

import (
	"errors"
	"fmt"

	"github.com.br/gibranct/ride/internal/infra/gateway"
	"github.com.br/gibranct/ride/internal/infra/repository"
)

type AcceptRideInput struct {
	DriverId string
	RideId   string
}

type AcceptRide struct {
	rideRepository repository.RideRepository
	accountGateway gateway.AccountGateway
}

func (ar *AcceptRide) Execute(input AcceptRideInput) error {
	account, err := ar.accountGateway.GetAccount(input.DriverId)
	if err != nil {
		return fmt.Errorf("account not found for id: %s", input.DriverId)
	}
	if !account.IsDriver {
		return errors.New("account must be a driver")
	}

	ride, err := ar.rideRepository.GetRideByID(input.RideId)
	if err != nil {
		return fmt.Errorf("ride not found: %s", err)
	}
	if err = ride.Accept(input.DriverId); err != nil {
		return err
	}
	return ar.rideRepository.UpdateRide(*ride)
}

func NewAcceptRideUseCase(accountGateway gateway.AccountGateway, rideRepo repository.RideRepository) *AcceptRide {
	return &AcceptRide{
		accountGateway: accountGateway,
		rideRepository: rideRepo,
	}
}
