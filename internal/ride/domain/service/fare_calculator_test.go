package service_test

import (
	"testing"
	"time"

	"github.com.br/gibranct/ride/internal/ride/domain/service"
	"github.com/stretchr/testify/assert"
)

func Test_NormalFareCalculator(t *testing.T) {
	distance := 0.0
	calculator := &service.NormalFareCalculator{}
	fare := calculator.Calculate(distance)
	assert.Equal(t, 0.0, fare)

	distance = 5.0
	fare = calculator.Calculate(distance)
	assert.Equal(t, 10.5, fare)

	distance = 10.0
	fare = calculator.Calculate(distance)
	assert.Equal(t, 21.0, fare)
}

func Test_OvernightFareCalculator(t *testing.T) {
	distance := 0.0
	calculator := &service.OvernightFareCalculator{}
	fare := calculator.Calculate(distance)
	assert.Equal(t, 0.0, fare)

	distance = 5.0
	fare = calculator.Calculate(distance)
	assert.Equal(t, 19.5, fare)

	distance = 10.0
	fare = calculator.Calculate(distance)
	assert.Equal(t, 39.0, fare)
}

func Test_SpecialDayFareCalculator(t *testing.T) {
	distance := 0.0
	calculator := &service.SpecialDayFareCalculator{}
	fare := calculator.Calculate(distance)
	assert.Equal(t, 0.0, fare)

	distance = 5.0
	fare = calculator.Calculate(distance)
	assert.Equal(t, 5.0, fare)

	distance = 10.0
	fare = calculator.Calculate(distance)
	assert.Equal(t, 10.0, fare)
}

func Test_NewFareCalculator(t *testing.T) {
	date := time.Date(2024, 1, 1, 22, 0, 0, 0, time.UTC)
	calculator := service.NewFareCalculator(date)
	assert.IsType(t, &service.SpecialDayFareCalculator{}, calculator)

	date = time.Date(2000, 1, 2, 22, 0, 0, 0, time.UTC)
	calculator = service.NewFareCalculator(date)
	assert.IsType(t, &service.OvernightFareCalculator{}, calculator)

	date = time.Date(2000, 2, 2, 5, 59, 0, 0, time.UTC)
	calculator = service.NewFareCalculator(date)
	assert.IsType(t, &service.OvernightFareCalculator{}, calculator)

	date = time.Date(2000, 2, 10, 10, 0, 0, 0, time.UTC)
	calculator = service.NewFareCalculator(date)
	assert.IsType(t, &service.NormalFareCalculator{}, calculator)
}
