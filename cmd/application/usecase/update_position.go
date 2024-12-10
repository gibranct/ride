package usecase

import (
	"fmt"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com.br/gibranct/ride/cmd/infra/repository"
)

type UpdatePositionInput struct {
	RideId string
	Lat    float64
	Long   float64
}

type UpdatePosition struct {
	rideRepository     repository.RideRepository
	positionRepository repository.PositionRepository
}

func (ar *UpdatePosition) Execute(input UpdatePositionInput) error {
	_, err := ar.rideRepository.GetRideByID(input.RideId)
	if err != nil {
		return fmt.Errorf("ride not found: %s", err)
	}
	newPosition, err := domain.CreatePosition(input.RideId, input.Lat, input.Long, nil)
	if err != nil {
		return err
	}
	return ar.positionRepository.SavePosition(*newPosition)
}

func NewUpdatePositionUseCase(
	rideRepo repository.RideRepository,
	positionRepo repository.PositionRepository,
) *UpdatePosition {
	return &UpdatePosition{
		rideRepository:     rideRepo,
		positionRepository: positionRepo,
	}
}
