package service_test

import (
	"testing"

	"github.com.br/gibranct/ride/internal/domain/service"
	"github.com.br/gibranct/ride/internal/domain/vo"
	"github.com/stretchr/testify/assert"
)

func Test_DistanceCalculator_Calculate_Equator(t *testing.T) {
	calculator := service.NewDistanceCalculator()
	from, err := vo.NewCoord(0, 0)
	assert.NoError(t, err)

	to, err := vo.NewCoord(0, 1)
	assert.NoError(t, err)

	distance := calculator.Calculate(*from, *to)
	expectedDistance := 111.0
	assert.InDelta(t, expectedDistance, distance, 1.0, "Distance calculation on equator should be approximately 111 km")
}

func Test_DistanceCalculator_Calculate_SameCoordinates(t *testing.T) {
	calculator := service.NewDistanceCalculator()
	coord, err := vo.NewCoord(40.7128, -74.0060)
	assert.NoError(t, err)
	distance := calculator.Calculate(*coord, *coord)
	assert.Equal(t, 0.0, distance, "Distance between same coordinates should be zero")
}

func Test_DistanceCalculator_Calculate_RoundedDistance(t *testing.T) {
	calculator := service.NewDistanceCalculator()
	from, err := vo.NewCoord(48.8566, 2.3522)
	assert.NoError(t, err)

	to, err := vo.NewCoord(51.5074, -0.1278)
	assert.NoError(t, err)

	distance := calculator.Calculate(*from, *to)
	expectedDistance := 344.0
	assert.Equal(t, expectedDistance, distance, "Distance should be rounded to the nearest whole kilometer")
}
