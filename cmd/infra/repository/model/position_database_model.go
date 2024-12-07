package model

import (
	"time"

	"github.com.br/gibranct/ride/cmd/domain"
)

type PositionDatabaseModel struct {
	PositionID string    `db:"position_id"`
	RideID     string    `db:"ride_id"`
	Lat        float64   `db:"lat"`
	Long       float64   `db:"long"`
	Date       time.Time `db:"date"`
}

func (e *PositionDatabaseModel) ToPosition() (*domain.Position, error) {
	return domain.NewPosition(
		e.PositionID,
		e.RideID,
		e.Lat,
		e.Long,
		e.Date,
	)
}
