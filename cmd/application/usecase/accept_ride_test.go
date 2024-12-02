package usecase_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com.br/gibranct/ride/cmd/application/usecase"
	di "github.com.br/gibranct/ride/cmd/infra/DI"
	"github.com/stretchr/testify/assert"
)

var (
	acceptRide = di.NewAcceptRide()
)

func Test_AcceptRide(t *testing.T) {
	inputSignupPassenger := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "secret",
	}
	outputSignUpPassenger, err := signUp.Execute(inputSignupPassenger)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, outputSignUpPassenger.AccountId)
	}

	inputSignupDriver := usecase.SignUpInput{
		Name:        "John Doe driver",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA8998",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}
	outputSignUpDriver, err := signUp.Execute(inputSignupDriver)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, outputSignUpDriver.AccountId)
	}
	rrInput := usecase.RequestRideInput{
		PassengerId: outputSignUpPassenger.AccountId,
		FromLat:     -27.584905257808835,
		FromLong:    -48.545022195325124,
		ToLat:       -27.496887588317275,
		ToLong:      -48.522234807851476,
	}
	outputRR, err := requestRide.Execute(rrInput)
	assert.NoError(t, err)
	inputAcceptRide := &usecase.AcceptRideInput{
		RideId:   outputRR.RideId,
		DriverId: outputSignUpDriver.AccountId,
	}
	err = acceptRide.Execute(inputAcceptRide)
	assert.NoError(t, err)

	outputGetRide, err := getRide.Execute(outputRR.RideId)
	if assert.NoError(t, err) {
		assert.Equal(t, "accepted", outputGetRide.Status)
		assert.Equal(t, outputSignUpDriver.AccountId, outputGetRide.DriverId)
	}
}
