package entity_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com.br/gibranct/account/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_CreateAccountWithoutID(t *testing.T) {
	name := "John Doe"
	email := fmt.Sprintf("john_%d@mail.com", rand.Int())
	cpf := "97456321558"
	carPlate := "AAA1234"
	password := "password"
	isPassenger := true
	isDriver := false
	newAccount, err := entity.CreateAccount(
		name,
		email,
		cpf,
		carPlate,
		password,
		isPassenger,
		isDriver,
	)
	assert.NoError(t, err)

	assert.Nil(t, err)
	assert.NotEmpty(t, newAccount.ID)
	assert.Equal(t, name, newAccount.GetName())
	assert.Equal(t, email, newAccount.GetEmail())
	assert.Equal(t, cpf, newAccount.GetCPF())
	assert.Empty(t, newAccount.GetCarPlate())
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(newAccount.GetPassword()), []byte(password)))
	assert.Equal(t, isPassenger, newAccount.IsPassenger)
	assert.Equal(t, isDriver, newAccount.IsDriver)
}

func Test_CreateAccountWithID(t *testing.T) {
	accountId := uuid.NewString()
	name := "John Doe"
	email := fmt.Sprintf("john_%d@mail.com", rand.Int())
	cpf := "97456321558"
	carPlate := "AAA1234"
	password := "password"
	isPassenger := true
	isDriver := false
	newAccount, err := entity.NewAccount(
		accountId,
		name,
		email,
		cpf,
		carPlate,
		password,
		isPassenger,
		isDriver,
	)

	assert.Nil(t, err)
	assert.Equal(t, accountId, newAccount.ID)
	assert.Equal(t, name, newAccount.GetName())
	assert.Equal(t, email, newAccount.GetEmail())
	assert.Equal(t, cpf, newAccount.GetCPF())
	assert.Empty(t, newAccount.GetCarPlate())
	assert.Equal(t, isPassenger, newAccount.IsPassenger)
	assert.Equal(t, isDriver, newAccount.IsDriver)
}

func Test_CreateAccountWithInvalidName(t *testing.T) {
	accountId := uuid.NewString()
	name := "John"
	email := fmt.Sprintf("john_%d@mail.com", rand.Int())
	cpf := "97456321558"
	carPlate := "AAA1234"
	password := "password"
	isPassenger := true
	isDriver := false
	newAccount, err := entity.NewAccount(
		accountId,
		name,
		email,
		cpf,
		carPlate,
		password,
		isPassenger,
		isDriver,
	)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid name", err.Error())
	assert.Nil(t, newAccount)
}

func Test_CreateAccountWithCarPlateNilIfIsPassenger(t *testing.T) {
	accountId := uuid.NewString()
	name := "John Doe"
	email := fmt.Sprintf("john_%d@mail.com", rand.Int())
	cpf := "97456321558"
	carPlate := ""
	password := "password"
	isPassenger := true
	isDriver := false
	newAccount, err := entity.NewAccount(
		accountId,
		name,
		email,
		cpf,
		carPlate,
		password,
		isPassenger,
		isDriver,
	)

	assert.Nil(t, err)
	assert.Empty(t, newAccount.GetCarPlate())
}
