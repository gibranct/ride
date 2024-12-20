package usecase

import (
	"fmt"
	"time"

	domain "github.com.br/gibranct/ride/internal/ride/domain/entity"
	"github.com.br/gibranct/ride/internal/ride/infra/repository"
)

type UpdatePositionInput struct {
	RideId string
	Lat    float64
	Long   float64
	Date   *time.Time
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
	newPosition, err := domain.CreatePosition(input.RideId, input.Lat, input.Long, input.Date)
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
