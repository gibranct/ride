package usecase_test

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com.br/gibranct/ride/internal/ride/application/usecase"
	di "github.com.br/gibranct/ride/internal/ride/infra/DI"
	"github.com.br/gibranct/ride/internal/ride/infra/gateway"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	acceptRide     = di.NewAcceptRide()
	accountGateway = gateway.NewAccountGateway()
)

func Test_AcceptRide(t *testing.T) {
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

	outputGetRide, err := getRide.Execute(outputRR.RideId)
	if assert.NoError(t, err) {
		assert.Equal(t, "accepted", outputGetRide.Status)
		assert.Equal(t, accountIdDriver, outputGetRide.DriverId)
	}
}

func Test_AcceptRideWhenAccountIsNotFound(t *testing.T) {
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
	accountId := uuid.NewString()

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
		DriverId: accountId,
	}
	err = acceptRide.Execute(inputAcceptRide)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "account not found")
}

func Test_AcceptRideWhenAccountIsNotADriver(t *testing.T) {
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
		DriverId: accountIdPassenger,
	}
	err = acceptRide.Execute(inputAcceptRide)

	assert.NotNil(t, err)
	assert.Equal(t, errors.New("account must be a driver"), err)
}
