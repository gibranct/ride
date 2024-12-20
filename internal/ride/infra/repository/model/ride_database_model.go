package model

import (
	"database/sql"
	"time"

	domain "github.com.br/gibranct/ride/internal/ride/domain/entity"
)

type RideDatabaseModel struct {
	RideID      string         `db:"ride_id"`
	PassengerID string         `db:"passenger_id"`
	DriverID    sql.NullString `db:"driver_id"`
	Status      string         `db:"status"`
	Fare        float64        `db:"fare"`
	Distance    float64        `db:"distance"`
	FromLat     float64        `db:"from_lat"`
	FromLong    float64        `db:"from_long"`
	ToLat       float64        `db:"to_lat"`
	ToLong      float64        `db:"to_long"`
	Date        time.Time      `db:"date"`
}

func (e *RideDatabaseModel) ToRide() (*domain.Ride, error) {
	return domain.NewRide(
		e.RideID,
		e.PassengerID,
		e.DriverID.String,
		e.FromLat,
		e.FromLong,
		e.ToLat,
		e.ToLong,
		e.Status,
		e.Date,
		e.Distance,
		e.Fare,
	)
}
