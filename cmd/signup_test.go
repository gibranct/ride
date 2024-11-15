package main

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignUpDriver(t *testing.T) {
	account := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}
	output, err := SignUp(account)

	if assert.NoError(t, err) {
		assert.NotEmpty(t, output.AccountId)
	}
}

func Test_SignUpDriverWithInvalidCarPlate(t *testing.T) {
	account := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA123",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}

	output, err := SignUp(account)

	if assert.Nil(t, output) {
		assert.Equal(t, err.Error(), "invalid car plate")
	}
}

func Test_SignUpPassenger(t *testing.T) {
	account := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "secret",
	}

	output, err := SignUp(account)

	if assert.NoError(t, err) {
		assert.NotEmpty(t, output.AccountId)
	}
}

func Test_SignUpPassengerWithInvalidCPF(t *testing.T) {
	account := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "9745632155",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}

	output, err := SignUp(account)

	if assert.Error(t, err) {
		assert.Nil(t, output)
		assert.Equal(t, err.Error(), "invalid cpf")
	}
}

func Test_SignUpPassengerWithInvalidEmail(t *testing.T) {
	account := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d_mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}

	output, err := SignUp(account)

	if assert.Error(t, err) {
		assert.Nil(t, output)
		assert.Equal(t, err.Error(), "invalid email")
	}
}

func Test_SignUpPassengerWithInvalidName(t *testing.T) {
	account := Account{
		Name:        "John",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}

	output, err := SignUp(account)

	if assert.Error(t, err) {
		assert.Nil(t, output)
		assert.Equal(t, err.Error(), "invalid name")
	}
}

func Test_SignUpDuplicatedPassenger(t *testing.T) {
	account := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}

	SignUp(account)
	output, err := SignUp(account)

	if assert.Error(t, err) {
		assert.Nil(t, output)
		assert.Equal(t, err.Error(), "duplicated account")
	}
}
