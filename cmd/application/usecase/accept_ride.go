package usecase

import (
	"errors"
	"fmt"

	"github.com.br/gibranct/ride/cmd/infra/repository"
)

type AcceptRideInput struct {
	DriverId string
	RideId   string
}

type AcceptRide struct {
	accountRepository repository.AccountRepository
	rideRepository    repository.RideRepository
}

func (ar *AcceptRide) Execute(input *AcceptRideInput) error {
	account, err := ar.accountRepository.GetAccountByID(input.DriverId)
	if err != nil {
		return fmt.Errorf("account not found: %s", err)
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

func NewAcceptRideUseCase(accountRepo repository.AccountRepository, rideRepo repository.RideRepository) *AcceptRide {
	return &AcceptRide{
		accountRepository: accountRepo,
		rideRepository:    rideRepo,
	}
}
