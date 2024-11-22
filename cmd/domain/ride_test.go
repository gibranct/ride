package domain_test

import (
	"testing"
	"time"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CreateRide(t *testing.T) {
	rideId := uuid.NewString()
	passengerId := uuid.NewString()
	fromLat := 89.0
	fromLong := 180.0
	toLat := 87.0
	toLong := 179.0
	status := "requested"
	date := time.Now()

	ride, err := domain.NewRide(rideId, passengerId, fromLat, fromLong, toLat, toLong, status, date)

	assert.Nil(t, err)
	assert.Equal(t, rideId, ride.GetRideId())
	assert.Equal(t, passengerId, ride.GetPassengerId())
	assert.Equal(t, fromLat, ride.GetFromCoord().GetLat())
	assert.Equal(t, fromLong, ride.GetFromCoord().GetLong())
	assert.Equal(t, toLat, ride.GetToCoord().GetLat())
	assert.Equal(t, toLong, ride.GetToCoord().GetLong())
	assert.Equal(t, status, ride.GetStatus())
	assert.Equal(t, date, *ride.GetDate())
}

func Test_CreateRideWithoutID(t *testing.T) {
	passengerId := uuid.NewString()
	fromLat := 89.0
	fromLong := 180.0
	toLat := 87.0
	toLong := 179.0
	status := "requested"

	ride, err := domain.CreateRide(passengerId, fromLat, fromLong, toLat, toLong)

	assert.Nil(t, err)
	assert.NotEmpty(t, ride.GetRideId())
	assert.Equal(t, passengerId, ride.GetPassengerId())
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
