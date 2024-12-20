package usecase

import (
	"errors"
	"fmt"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com.br/gibranct/ride/cmd/infra/repository"
)

type RequestRide struct {
	accountRepository repository.AccountRepository
	rideRepository    repository.RideRepository
}

type RequestRideInput struct {
	PassengerId string
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
}

type RequestRideOutput struct {
	RideId string
}

func (rr *RequestRide) Execute(input RequestRideInput) (*RequestRideOutput, error) {
	account, err := rr.accountRepository.GetAccountByID(input.PassengerId)
	if err != nil {
		return nil, fmt.Errorf("account %s does not exist", input.PassengerId)
	}
	if !account.IsPassenger {
		return nil, errors.New("account must be from a passenger")
	}
	ride, err := domain.CreateRide(input.PassengerId, input.FromLat, input.FromLong, input.ToLat, input.ToLong)
	if err != nil {
		return nil, err
	}
	if err = rr.rideRepository.SaveRide(*ride); err != nil {
		return nil, err
	}

	return &RequestRideOutput{
		RideId: ride.GetRideId(),
	}, nil
}

func NewRequestRideUseCase(accountRepo repository.AccountRepository, rideRepo repository.RideRepository) *RequestRide {
	return &RequestRide{
		accountRepository: accountRepo,
		rideRepository:    rideRepo,
	}
}
