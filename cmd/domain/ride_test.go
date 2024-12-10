package domain_test

import (
	"testing"
	"time"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Test_CreateRide(t *testing.T) {
	rideId := uuid.NewString()
	passengerId := uuid.NewString()
	driverId := uuid.NewString()
	fromLat := 89.0
	fromLong := 180.0
	toLat := 87.0
	toLong := 179.0
	status := "requested"
	date := time.Now()

	ride, err := domain.NewRide(rideId, passengerId, driverId, fromLat, fromLong, toLat, toLong, status, date)

	assert.Nil(t, err)
	assert.Equal(t, rideId, ride.GetRideId())
	assert.Equal(t, passengerId, ride.GetPassengerId())
	assert.Equal(t, driverId, ride.GetDriverId())
	assert.Equal(t, fromLat, ride.GetFromCoord().GetLat())
	assert.Equal(t, fromLong, ride.GetFromCoord().GetLong())
	assert.Equal(t, toLat, ride.GetToCoord().GetLat())
	assert.Equal(t, toLong, ride.GetToCoord().GetLong())
	assert.Equal(t, status, ride.GetStatus())
	assert.Equal(t, date, *ride.GetDate())
}

func Test_CreateRideWithoutID(t *testing.T) {
	passengerId := uuid.NewString()
	driverId := ""
	fromLat := 89.0
	fromLong := 180.0
	toLat := 87.0
	toLong := 179.0
	status := "requested"

	ride, err := domain.CreateRide(passengerId, fromLat, fromLong, toLat, toLong)

	assert.Nil(t, err)
	assert.NotEmpty(t, ride.GetRideId())
	assert.Equal(t, passengerId, ride.GetPassengerId())
	assert.Equal(t, driverId, ride.GetDriverId())
	assert.Equal(t, fromLat, ride.GetFromCoord().GetLat())
	assert.Equal(t, fromLong, ride.GetFromCoord().GetLong())
	assert.Equal(t, toLat, ride.GetToCoord().GetLat())
	assert.Equal(t, toLong, ride.GetToCoord().GetLong())
	assert.Equal(t, status, ride.GetStatus())
	assert.NotNil(t, ride.GetDate())
}

func Test_CreateRideWithoutInvalidLat(t *testing.T) {
	passengerId := uuid.NewString()
	fromLat := 91.0
	fromLong := 180.0
	toLat := 87.0
	toLong := 179.0

	ride, err := domain.CreateRide(passengerId, fromLat, fromLong, toLat, toLong)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid latitude", err.Error())
	assert.Nil(t, ride)
}

func Test_RideFinishWithOnePosition(t *testing.T) {
	rideId := uuid.NewString()
	passengerId := uuid.NewString()
	driverId := uuid.NewString()
	fromLat := 0.0
	fromLong := 0.0
	toLat := 0.0
	toLong := 0.0
	status := "in_progress"
	date := time.Now()

	ride, err := domain.NewRide(rideId, passengerId, driverId, fromLat, fromLong, toLat, toLong, status, date)
	assert.NoError(t, err)

	p1, err := domain.CreatePosition(rideId, 0.0, 0.0, &date)
	assert.NoError(t, err)

	positions := []domain.Position{*p1}

	ride.Finish(positions)

	assert.Equal(t, 0.0, ride.GetDistance(positions))
	assert.Equal(t, 0.0, ride.GetFare())
}

func Test_RideFinishWithEmptyPositions(t *testing.T) {
	rideId := uuid.NewString()
	passengerId := uuid.NewString()
	driverId := uuid.NewString()
	fromLat := 89.0
	fromLong := 180.0
	toLat := 87.0
	toLong := 179.0
	status := "in_progress"
	date := time.Now()

	ride, err := domain.NewRide(rideId, passengerId, driverId, fromLat, fromLong, toLat, toLong, status, date)
	assert.NoError(t, err)

	ride.Finish([]domain.Position{})

	assert.Equal(t, 0.0, ride.GetDistance([]domain.Position{}))
	assert.Equal(t, 0.0, ride.GetFare())
}
