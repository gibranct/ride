package usecase_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com.br/gibranct/ride/cmd/application/usecase"
	di "github.com.br/gibranct/ride/cmd/infra/DI"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	signUp      = di.NewSignUp()
	requestRide = di.NewRequestRide()
	getRide     = di.NewGetRide()
	getAccount  = di.NewGetAccount()
)

func Test_RequestRide(t *testing.T) {
	signupInput := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "secret123",
	}
	outputSignUp, err := signUp.Execute(signupInput)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, outputSignUp.AccountId)
	}

	rrInput := usecase.RequestRideInput{
		PassengerId: outputSignUp.AccountId,
		FromLat:     -27.584905257808835,
		FromLong:    -48.545022195325124,
		ToLat:       -27.496887588317275,
		ToLong:      -48.522234807851476,
	}
	outputRR, err := requestRide.Execute(rrInput)

	if assert.NoError(t, err) {
		assert.NotEmpty(t, outputRR.RideId)
		outputGR, err := getRide.Execute(outputRR.RideId)
		if assert.NoError(t, err) {
			assert.Equal(t, outputRR.RideId, outputGR.RideId)
			assert.Equal(t, rrInput.PassengerId, outputGR.PassengerId)
			assert.Equal(t, rrInput.FromLat, outputGR.FromLat)
			assert.Equal(t, rrInput.FromLong, outputGR.FromLong)
			assert.Equal(t, rrInput.ToLat, outputGR.ToLat)
			assert.Equal(t, rrInput.ToLong, outputGR.ToLong)
		}
	}
}

func Test_RequestRideForDriver(t *testing.T) {
	signupInput := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA5887",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret123",
	}
	outputSignUp, err := signUp.Execute(signupInput)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, outputSignUp.AccountId)
	}

	rrInput := usecase.RequestRideInput{
		PassengerId: outputSignUp.AccountId,
		FromLat:     -27.584905257808835,
		FromLong:    -48.545022195325124,
		ToLat:       -27.496887588317275,
		ToLong:      -48.522234807851476,
	}
	outputRR, err := requestRide.Execute(rrInput)
	assert.Nil(t, outputRR)
	assert.Equal(t, err.Error(), "account must be from a passenger")
}

func Test_RequestRideWhenAccountDoesNotExist(t *testing.T) {
	accountId := uuid.NewString()
	rrInput := usecase.RequestRideInput{
		PassengerId: accountId,
		FromLat:     -27.584905257808835,
		FromLong:    -48.545022195325124,
		ToLat:       -27.496887588317275,
		ToLong:      -48.522234807851476,
	}
	outputRR, err := requestRide.Execute(rrInput)
	assert.Nil(t, outputRR)
	assert.Equal(t, err.Error(), fmt.Sprintf("account %s does not exist", accountId))
}
