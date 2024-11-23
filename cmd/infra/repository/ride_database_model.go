package repository

import (
	"time"

	"github.com.br/gibranct/ride/cmd/domain"
)

type RideDatabaseModel struct {
	RideID      string
	PassengerID string
	DriverID    string
	Status      string
	Fare        float32
	Distance    float32
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
	Date        time.Time
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
