package repository

import (
	"context"

	domain "github.com.br/gibranct/ride/internal/ride/domain/entity"
	"github.com.br/gibranct/ride/internal/ride/infra/database"
	"github.com.br/gibranct/ride/internal/ride/infra/repository/model"
)

type RideRepository interface {
	GetRideByID(id string) (*domain.Ride, error)
	SaveRide(ride domain.Ride) error
	UpdateRide(ride domain.Ride) error
	HasActiveRideByPassengerId(passengerId string) (bool, error)
}

type RideRepositoryDatabase struct {
	connection database.DatabaseConnection
}

func (repo RideRepositoryDatabase) GetRideByID(id string) (*domain.Ride, error) {
	var rideModel model.RideDatabaseModel
	query := "select * from gct.ride where ride_id = $1"
	err := repo.connection.QueryWithContext(context.Background(), &rideModel, query, id)
	if err != nil {
		return nil, err
	}

	return rideModel.ToRide()
}

func (repo RideRepositoryDatabase) SaveRide(ride domain.Ride) error {
	saveQuery := "insert into gct.ride (ride_id, passenger_id, from_lat, from_long, to_lat, to_long, status, date, distance, fare) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	args := []any{
		ride.GetRideId(), ride.GetPassengerId(),
		ride.GetFromCoord().GetLat(), ride.GetFromCoord().GetLong(),
		ride.GetToCoord().GetLat(), ride.GetToCoord().GetLong(),
		ride.GetStatus(), ride.GetDate(),
		ride.GetDistance(), ride.GetFare(),
	}
	return repo.connection.ExecContext(context.Background(), saveQuery, args...)
}

func (repo RideRepositoryDatabase) UpdateRide(ride domain.Ride) error {
	updateQuery := "update gct.ride set status = $1, driver_id = $2, distance = $3, fare = $4 where ride_id = $5"
	args := []any{
		ride.GetStatus(), ride.GetDriverId(), ride.GetDistance(), ride.GetFare(), ride.GetRideId(),
	}
	return repo.connection.ExecContext(context.Background(), updateQuery, args...)
}

func (repo RideRepositoryDatabase) HasActiveRideByPassengerId(passengerId string) (bool, error) {
	var count int
	query := "select count(*) from gct.ride where passenger_id = $1 and status not in ('completed', 'cancelled')"
	err := repo.connection.QueryWithContext(context.Background(), &count, query, passengerId)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func NewRideRepository(connection database.DatabaseConnection) *RideRepositoryDatabase {
	return &RideRepositoryDatabase{
		connection: connection,
	}
}
