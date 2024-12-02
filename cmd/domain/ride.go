package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ride struct {
	rideId      string
	passengerId string
	driverId    string
	from        *Coord
	to          *Coord
	status      RideStatus
	date        *time.Time
}

func NewRide(
	rideId, passengerId, driverId string, fromLat, fromLong, toLat, toLong float64, status string, date time.Time,
) (*Ride, error) {
	fromCoord, err := NewCoord(fromLat, fromLong)
	if err != nil {
		return nil, err
	}

	toCoord, err := NewCoord(toLat, toLong)
	if err != nil {
		return nil, err
	}

	ride := &Ride{
		rideId:      rideId,
		passengerId: passengerId,
		driverId:    driverId,
		from:        fromCoord,
		to:          toCoord,
		date:        &date,
	}

	rideStatus, err := NewRideStatus(status, ride)
	ride.SetStatus(rideStatus)
	if err != nil {
		return nil, err
	}

	return ride, nil
}

func CreateRide(
	passengerId string, fromLat, fromLong, toLat, toLong float64,
) (*Ride, error) {
	rideId := uuid.NewString()
	status := "requested"
	now := time.Now()
	return NewRide(
		rideId, passengerId, "", fromLat, fromLong, toLat, toLong, status, now,
	)
}

func (r *Ride) GetRideId() string {
	return r.rideId
}

func (r *Ride) GetPassengerId() string {
	return r.passengerId
}

func (r *Ride) GetDriverId() string {
	return r.driverId
}

func (r *Ride) GetFromCoord() *Coord {
	return r.from
}

func (r *Ride) GetToCoord() *Coord {
	return r.to
}

func (r *Ride) GetStatus() string {
	return r.status.GetValue()
}

func (r *Ride) GetDate() *time.Time {
	return r.date
}

func (r *Ride) SetStatus(status RideStatus) {
	r.status = status
}

func (r *Ride) Accept(driverId string) error {
	if err := r.status.Accept(); err != nil {
		return err
	}
	r.driverId = driverId
	return nil
}
