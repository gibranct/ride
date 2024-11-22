package repository

import (
	"context"
	"fmt"
	"os"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com/jackc/pgx/v5"
)

type RideRepository interface {
	GetRideByID(id string) (*domain.Ride, error)
	SaveRide(account domain.Ride) error
}

type RideRepositoryDatabase struct{}

func (repo RideRepositoryDatabase) GetRideByID(id string) (*domain.Ride, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	var rideModel RideDatabaseModel
	conn.QueryRow(context.Background(), "select ride_id, passenger_id, from_lat, from_long, to_lat, to_long, status, date from gct.ride where ride_id = $1", id).Scan(
		&rideModel.RideID, &rideModel.PassengerID, &rideModel.FromLat, &rideModel.FromLong, &rideModel.ToLat, &rideModel.ToLong, &rideModel.Status, &rideModel.Date,
	)

	return rideModel.ToRide()
}

func (repo RideRepositoryDatabase) SaveRide(ride domain.Ride) error {
	saveQuery := "insert into gct.ride (ride_id, passenger_id, from_lat, from_long, to_lat, to_long, status, date) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	args := []any{
		ride.GetRideId(), ride.GetPassengerId(), ride.GetFromCoord().GetLat(), ride.GetFromCoord().GetLong(), ride.GetToCoord().GetLat(), ride.GetToCoord().GetLong(), ride.GetStatus(), ride.GetDate(),
	}
	_, err = conn.Exec(context.Background(), saveQuery, args...)

	return err
}

func NewRideRepository() RideRepository {
	return &RideRepositoryDatabase{}
}
