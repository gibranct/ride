package repository

import (
	"context"

	domain "github.com.br/gibranct/ride/internal/ride/domain/entity"
	"github.com.br/gibranct/ride/internal/ride/infra/database"
	"github.com.br/gibranct/ride/internal/ride/infra/repository/model"
)

type PositionRepository interface {
	SavePosition(position domain.Position) error
	GetPositionsByRideId(rideId string) ([]domain.Position, error)
}

type PositionRepositoryDatabase struct {
	connection database.DatabaseConnection
}

func (repo *PositionRepositoryDatabase) SavePosition(position domain.Position) error {
	query := "insert into gct.position (position_id, ride_id, lat, long, date) values ($1, $2, $3, $4, $5)"
	position.GetCoord()
	args := []any{
		position.GetPositionId(), position.GetRideId(), position.GetCoord().GetLat(), position.GetCoord().GetLong(), position.Date,
	}
	return repo.connection.ExecContext(context.Background(), query, args...)
}

func (repo *PositionRepositoryDatabase) GetPositionsByRideId(rideId string) ([]domain.Position, error) {
	modelPositions := make([]model.PositionDatabaseModel, 0)
	query := "select position_id, ride_id, lat, long, date from gct.position where ride_id = $1"
	err := repo.connection.SelectContext(context.Background(), &modelPositions, query, rideId)
	if err != nil {
		return nil, err
	}
	positions := make([]domain.Position, len(modelPositions))
	for i, modelPosition := range modelPositions {
		pos, err := modelPosition.ToPosition()
		if err != nil {
			return nil, err
		}
		positions[i] = *pos
	}
	return positions, nil
}

func NewPositionRepository(conn database.DatabaseConnection) *PositionRepositoryDatabase {
	return &PositionRepositoryDatabase{
		connection: conn,
	}
}
