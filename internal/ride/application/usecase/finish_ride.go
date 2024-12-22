package usecase

import (
	"encoding/json"
	"fmt"

	"github.com.br/gibranct/ride/internal/ride/domain/event"
	"github.com.br/gibranct/ride/internal/ride/infra/queue"
	"github.com.br/gibranct/ride/internal/ride/infra/repository"
)

type FinishRideInput struct {
	RideId string
}

type FinishRide struct {
	rideRepository     repository.RideRepository
	positionRepository repository.PositionRepository
	queue              queue.Queue
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
	err = ar.rideRepository.UpdateRide(*ride)
	if err != nil {
		return err
	}
	event := event.NewFinishRideEvent(ride.GetRideId(), ride.GetFare())
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return ar.queue.Publish("rideCompleted", bytes)
}

func NewFinishRideUseCase(
	rideRepo repository.RideRepository,
	positionRepo repository.PositionRepository,
	queue queue.Queue,
) *FinishRide {
	return &FinishRide{
		rideRepository:     rideRepo,
		positionRepository: positionRepo,
		queue:              queue,
	}
}
