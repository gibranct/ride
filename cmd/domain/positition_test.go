package domain_test

import (
	"testing"
	"time"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com/stretchr/testify/assert"
)

func Test_NewPosition(t *testing.T) {
	positionId := "position123"
	rideId := "ride123"
	lat := -23.5489
	long := -46.6388
	now := time.Now()

	position, err := domain.NewPosition(positionId, rideId, lat, long, now)
	assert.Nil(t, err)

	assert.Equal(t, positionId, position.GetPositionId())
	assert.Equal(t, rideId, position.GetRideId())
	assert.Equal(t, lat, position.GetCoord().GetLat())
	assert.Equal(t, long, position.GetCoord().GetLong())
}

func Test_CreatePosition(t *testing.T) {
	rideId := "ride123"
	lat := -23.5489
	long := -46.6388

	position, err := domain.CreatePosition(rideId, lat, long)
	assert.Nil(t, err)

	assert.NotEmpty(t, position.GetPositionId())
	assert.Equal(t, rideId, position.GetRideId())
	assert.Equal(t, lat, position.GetCoord().GetLat())
	assert.Equal(t, long, position.GetCoord().GetLong())
}
