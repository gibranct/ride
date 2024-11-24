package repository

import (
	"context"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com.br/gibranct/ride/cmd/infra/database"
)

type RideRepository interface {
	GetRideByID(id string) (*domain.Ride, error)
	SaveRide(ride domain.Ride) error
}

type RideRepositoryDatabase struct {
	connection database.DatabaseConnection
}

func (repo RideRepositoryDatabase) GetRideByID(id string) (*domain.Ride, error) {
	var rideModel RideDatabaseModel
	query := "select ride_id, passenger_id, from_lat, from_long, to_lat, to_long, status, date from gct.ride where ride_id = $1"
	err := repo.connection.QueryWithContext(context.Background(), &rideModel, query, id)
	if err != nil {
		return nil, err
	}

	return rideModel.ToRide()
}

func (repo RideRepositoryDatabase) SaveRide(ride domain.Ride) error {
	saveQuery := "insert into gct.ride (ride_id, passenger_id, from_lat, from_long, to_lat, to_long, status, date) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	args := []any{
		ride.GetRideId(), ride.GetPassengerId(), ride.GetFromCoord().GetLat(), ride.GetFromCoord().GetLong(), ride.GetToCoord().GetLat(), ride.GetToCoord().GetLong(), ride.GetStatus(), ride.GetDate(),
	}
	return repo.connection.ExecContext(context.Background(), saveQuery, args...)
}

func NewRideRepository(connection database.DatabaseConnection) RideRepository {
	return &RideRepositoryDatabase{
		connection: connection,
	}
}
