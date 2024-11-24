package repository

import (
	"time"

	"github.com.br/gibranct/ride/cmd/domain"
)

type RideDatabaseModel struct {
	RideID      string    `db:"ride_id"`
	PassengerID string    `db:"passenger_id"`
	DriverID    string    `db:"driver_id"`
	Status      string    `db:"status"`
	Fare        float32   `db:"fare"`
	Distance    float32   `db:"distance"`
	FromLat     float64   `db:"from_lat"`
	FromLong    float64   `db:"from_long"`
	ToLat       float64   `db:"to_lat"`
	ToLong      float64   `db:"to_long"`
	Date        time.Time `db:"date"`
}

func (e *RideDatabaseModel) ToRide() (*domain.Ride, error) {
	return domain.NewRide(
		e.RideID,
		e.PassengerID,
		e.FromLat,
		e.FromLong,
		e.ToLat,
		e.ToLong,
		e.Status,
		e.Date,
	)
}
