package usecase

import (
	"errors"
	"fmt"

	domain "github.com.br/gibranct/ride/internal/ride/domain/entity"
	"github.com.br/gibranct/ride/internal/ride/infra/gateway"
	"github.com.br/gibranct/ride/internal/ride/infra/repository"
)

type RequestRide struct {
	accountGateway gateway.AccountGateway
	rideRepository repository.RideRepository
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
	account, err := rr.accountGateway.GetAccount(input.PassengerId)
	if err != nil {
		return nil, fmt.Errorf("account %s does not exist", input.PassengerId)
	}
	if !account.IsPassenger {
		return nil, errors.New("account must be from a passenger")
	}
	passengerHasActiveRide, err := rr.rideRepository.HasActiveRideByPassengerId(account.ID)
	if err != nil {
		return nil, err
	}
	if passengerHasActiveRide {
		return nil, errors.New("you already have an active ride")
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

func NewRequestRideUseCase(accountRepo gateway.AccountGateway, rideRepo repository.RideRepository) *RequestRide {
	return &RequestRide{
		accountGateway: accountRepo,
		rideRepository: rideRepo,
	}
}
