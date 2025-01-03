package usecase_test

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com.br/gibranct/ride/internal/application/usecase"
	domain "github.com.br/gibranct/ride/internal/domain/entity"
	di "github.com.br/gibranct/ride/internal/infra/DI"
	"github.com.br/gibranct/ride/internal/infra/gateway"
	"github.com/stretchr/testify/assert"
)

var (
	finishRide = di.NewFinishRide()
)

func Test_FinishRideNormalHours(t *testing.T) {
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

	date := time.Date(2022, 02, 15, 6, 0, 0, 0, time.UTC)

	inputUpdatePosition1 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.584905257808835,
		Long:   -48.545022195325124,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition1)
	assert.NoError(t, err)

	inputUpdatePosition2 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.496887588317275,
		Long:   -48.522234807851476,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition2)
	assert.NoError(t, err)

	inputUpdatePosition3 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.584905257808835,
		Long:   -48.545022195325124,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition3)
	assert.NoError(t, err)

	inputUpdatePosition4 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.496887588317275,
		Long:   -48.522234807851476,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition4)
	assert.NoError(t, err)

	inputFinishRide := usecase.FinishRideInput{
		RideId: outputRR.RideId,
	}
	finishRide.Execute(inputFinishRide)
	outputGetRide, err := getRide.Execute(outputRR.RideId)
	assert.NoError(t, err)

	assert.Equal(t, float64(30), outputGetRide.Distance)
	assert.Equal(t, float64(63), outputGetRide.Fare)
	assert.Equal(t, domain.COMPLETED_RIDE_STATUS, outputGetRide.Status)
}

func Test_FinishRideOvernight(t *testing.T) {
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

	date := time.Date(2022, 02, 15, 23, 0, 0, 0, time.UTC)

	inputUpdatePosition1 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.584905257808835,
		Long:   -48.545022195325124,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition1)
	assert.NoError(t, err)

	inputUpdatePosition2 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.496887588317275,
		Long:   -48.522234807851476,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition2)
	assert.NoError(t, err)

	inputUpdatePosition3 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.584905257808835,
		Long:   -48.545022195325124,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition3)
	assert.NoError(t, err)

	inputUpdatePosition4 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.496887588317275,
		Long:   -48.522234807851476,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition4)
	assert.NoError(t, err)

	inputFinishRide := usecase.FinishRideInput{
		RideId: outputRR.RideId,
	}
	finishRide.Execute(inputFinishRide)
	outputGetRide, err := getRide.Execute(outputRR.RideId)
	assert.NoError(t, err)

	assert.Equal(t, float64(30), outputGetRide.Distance)
	assert.Equal(t, float64(117), outputGetRide.Fare)
	assert.Equal(t, domain.COMPLETED_RIDE_STATUS, outputGetRide.Status)
}

func Test_FinishRideSpecialDay(t *testing.T) {
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

	date := time.Date(2022, 02, 1, 23, 0, 0, 0, time.UTC)

	inputUpdatePosition1 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.584905257808835,
		Long:   -48.545022195325124,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition1)
	assert.NoError(t, err)

	inputUpdatePosition2 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.496887588317275,
		Long:   -48.522234807851476,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition2)
	assert.NoError(t, err)

	inputUpdatePosition3 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.584905257808835,
		Long:   -48.545022195325124,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition3)
	assert.NoError(t, err)

	inputUpdatePosition4 := usecase.UpdatePositionInput{
		RideId: outputRR.RideId,
		Lat:    -27.496887588317275,
		Long:   -48.522234807851476,
		Date:   &date,
	}
	err = updatePosition.Execute(inputUpdatePosition4)
	assert.NoError(t, err)

	inputFinishRide := usecase.FinishRideInput{
		RideId: outputRR.RideId,
	}
	finishRide.Execute(inputFinishRide)
	outputGetRide, err := getRide.Execute(outputRR.RideId)
	assert.NoError(t, err)

	assert.Equal(t, float64(30), outputGetRide.Distance)
	assert.Equal(t, float64(30), outputGetRide.Fare)
	assert.Equal(t, domain.COMPLETED_RIDE_STATUS, outputGetRide.Status)
}
