package usecase_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com.br/gibranct/ride/internal/account/application/usecase"
	di "github.com.br/gibranct/ride/internal/account/infra/DI"
	"github.com/stretchr/testify/assert"
)

var (
	signUp     = di.NewSignUp()
	getAccount = di.NewGetAccount()
)

func Test_SignUpDriver(t *testing.T) {
	account := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret123",
	}
	output, err := signUp.Execute(account)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, output.AccountId)
		accountOutput, err := getAccount.Execute(output.AccountId)
		if assert.NoError(t, err) {
			assert.Equal(t, output.AccountId, output.AccountId)
			assert.Equal(t, account.Name, accountOutput.Name)
			assert.Equal(t, account.Email, accountOutput.Email)
			assert.Equal(t, account.CPF, accountOutput.CPF)
			assert.Equal(t, account.CarPlate, accountOutput.CarPlate)
			assert.True(t, accountOutput.IsDriver)
			assert.False(t, accountOutput.IsPassenger)
		}
	}
}

func Test_SignUpPassenger(t *testing.T) {
	account := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "secret123",
	}
	output, err := signUp.Execute(account)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, output.AccountId)
		accountOutput, err := getAccount.Execute(output.AccountId)
		if assert.NoError(t, err) {
			assert.Equal(t, output.AccountId, output.AccountId)
			assert.Equal(t, account.Name, accountOutput.Name)
			assert.Equal(t, account.Email, accountOutput.Email)
			assert.Equal(t, account.CPF, accountOutput.CPF)
			assert.Equal(t, account.CarPlate, accountOutput.CarPlate)
			assert.True(t, accountOutput.IsPassenger)
			assert.False(t, accountOutput.IsDriver)
		}
	}
}

func Test_SignUpPassengerWithInvalidEmail(t *testing.T) {
	account := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d_mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret123",
	}

	output, err := signUp.Execute(account)

	if assert.Error(t, err) {
		assert.Nil(t, output)
		assert.Equal(t, err.Error(), "invalid email")
	}
}

func Test_SignUpDuplicatedPassenger(t *testing.T) {
	account := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret123",
	}

	signUp.Execute(account)
	output, err := signUp.Execute(account)

	if assert.Error(t, err) {
		assert.Nil(t, output)
		assert.Equal(t, err.Error(), "duplicated account")
	}
}
