package vo_test

import (
	"testing"

	"github.com.br/gibranct/ride/internal/ride/domain/vo"
	"github.com/stretchr/testify/assert"
)

func Test_CreateCoord(t *testing.T) {
	points := [][]float64{{-90, -180}, {90, 180}, {-89, -179}, {89, 179}}
	for _, point := range points {
		coord, err := vo.NewCoord(point[0], point[1])
		assert.Nil(t, err)
		assert.NotNil(t, coord)
		assert.Equal(t, point[0], coord.GetLat())
		assert.Equal(t, point[1], coord.GetLong())
	}
}

func Test_CreateCoordWithInvalidLat(t *testing.T) {
	invalidLats := []float64{-91, 91}
	long := 180.0
	for _, lat := range invalidLats {
		coord, err := vo.NewCoord(lat, long)
		assert.NotNil(t, err)
		assert.Equal(t, "invalid latitude", err.Error())
		assert.Nil(t, coord)
	}
}

func Test_CreateCoordWithInvalidLong(t *testing.T) {
	invalidLongs := []float64{-181, 181}
	lat := 90.0
	for _, long := range invalidLongs {
		coord, err := vo.NewCoord(lat, long)
		assert.NotNil(t, err)
		assert.Equal(t, "invalid longitude", err.Error())
		assert.Nil(t, coord)
	}
}
