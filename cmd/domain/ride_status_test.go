package domain_test

import (
	"errors"
	"testing"
	"time"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_RequestedRideStatus(t *testing.T) {
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
	assert.NoError(t, err)
	assert.Equal(t, status, ride.GetStatus())

	ridStatus := domain.NewRequestedStatus(ride)

	assert.Equal(t, domain.REQUESTED_RIDE_STATUS, ridStatus.GetValue())
	assert.Equal(t, errors.New("invalid status"), ridStatus.Start())
	assert.NoError(t, ride.Accept(driverId))
	assert.Equal(t, domain.ACCEPTED_RIDE_STATUS, ride.GetStatus())
}

func Test_AcceptedRideStatus(t *testing.T) {
	rideId := uuid.NewString()
	passengerId := uuid.NewString()
	driverId := uuid.NewString()
	fromLat := 89.0
	fromLong := 180.0
	toLat := 87.0
	toLong := 179.0
	status := "accepted"
	date := time.Now()

	ride, err := domain.NewRide(rideId, passengerId, driverId, fromLat, fromLong, toLat, toLong, status, date)
	assert.NoError(t, err)
	assert.Equal(t, status, ride.GetStatus())

	ridStatus := domain.NewAcceptedStatus(ride)

	assert.Equal(t, domain.ACCEPTED_RIDE_STATUS, ridStatus.GetValue())
	assert.Equal(t, domain.ACCEPTED_RIDE_STATUS, ride.GetStatus())
	assert.Equal(t, errors.New("invalid status"), ridStatus.Accept())
	assert.Equal(t, errors.New("invalid status"), ridStatus.Request())
	assert.Nil(t, ridStatus.Start())
}

func Test_InProgressRideStatus(t *testing.T) {
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
	assert.Equal(t, status, ride.GetStatus())

	ridStatus := domain.NewInProgressStatus(ride)

	assert.Equal(t, domain.IN_PROGRESS_RIDE_STATUS, ridStatus.GetValue())
	assert.Equal(t, errors.New("invalid status"), ridStatus.Request())
	assert.Equal(t, errors.New("invalid status"), ridStatus.Accept())
	assert.Equal(t, errors.New("invalid status"), ridStatus.Start())
	assert.Equal(t, domain.IN_PROGRESS_RIDE_STATUS, ride.GetStatus())
}
