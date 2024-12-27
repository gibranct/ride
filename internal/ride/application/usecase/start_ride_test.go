package usecase_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com.br/gibranct/ride/internal/ride/application/usecase"
	di "github.com.br/gibranct/ride/internal/ride/infra/DI"
	"github.com.br/gibranct/ride/internal/ride/infra/gateway"
	"github.com/stretchr/testify/assert"
)

var (
	startRide = di.NewStartRide()
)

func Test_StartRide(t *testing.T) {
	inputSignupPassenger := gateway.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "secret123",
	}
	accountIdPassenger, err := accountGateway.SignUp(inputSignupPassenger)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, accountIdPassenger)
	}

	inputSignupDriver := gateway.SignUpInput{
		Name:        "John Doe driver",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA8998",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret123",
	}
	accountIdDriver, err := accountGateway.SignUp(inputSignupDriver)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, accountIdDriver)
	}
	rrInput := usecase.RequestRideInput{
		PassengerId: accountIdPassenger,
		FromLat:     -27.584905257808835,
		FromLong:    -48.545022195325124,
		ToLat:       -27.496887588317275,
		ToLong:      -48.522234807851476,
	}
	outputRR, err := requestRide.Execute(rrInput)
	assert.NoError(t, err)
	inputAcceptRide := usecase.AcceptRideInput{
		RideId:   outputRR.RideId,
		DriverId: accountIdDriver,
	}
	err = acceptRide.Execute(inputAcceptRide)
	assert.NoError(t, err)

	inputStartRide := usecase.StartRideInput{
		RideId: outputRR.RideId,
	}

	err = startRide.Execute(inputStartRide)
	assert.NoError(t, err)

	outputGetRide, err := getRide.Execute(outputRR.RideId)
	if assert.NoError(t, err) {
		assert.Equal(t, "in_progress", outputGetRide.Status)
	}
}
