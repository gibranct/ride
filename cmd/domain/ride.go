package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ride struct {
	rideId      string
	passengerId string
	from        *Coord
	to          *Coord
	status      string
	date        *time.Time
}

func NewRide(
	rideId, passengerId string, fromLat, fromLong, toLat, toLong float64, status string, date time.Time,
) (*Ride, error) {
	fromCoord, err := NewCoord(fromLat, fromLong)
	if err != nil {
		return nil, err
	}

	toCoord, err := NewCoord(toLat, toLong)
	if err != nil {
		return nil, err
	}

	return &Ride{
		rideId:      rideId,
		passengerId: passengerId,
		from:        fromCoord,
		to:          toCoord,
		status:      status,
		date:        &date,
	}, nil
}

func CreateRide(
	passengerId string, fromLat, fromLong, toLat, toLong float64,
) (*Ride, error) {
	rideId := uuid.NewString()
	status := "requested"
	now := time.Now()
	return NewRide(
		rideId, passengerId, fromLat, fromLong, toLat, toLong, status, now,
	)
}

func (r *Ride) GetRideId() string {
	return r.rideId
}

func (r *Ride) GetPassengerId() string {
	return r.passengerId
}

func (r *Ride) GetFromCoord() *Coord {
	return r.from
}

func (r *Ride) GetToCoord() *Coord {
	return r.to
}

func (r *Ride) GetStatus() string {
	return r.status
}

func (r *Ride) GetDate() *time.Time {
	return r.date
}
