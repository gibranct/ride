package test

import (
	"fmt"
	"log"
	"math/rand/v2"
	"testing"

	"github.com.br/gibranct/account/internal/domain/entity"
	di "github.com.br/gibranct/account/internal/infra/DI"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SaveAccount(t *testing.T) {
	repo := di.NewAccountPostgresRepository()

	account := getAccount()

	assert.NoError(t, repo.SaveAccount(*account))

	savedAccount, err := repo.GetAccountByEmail(account.GetEmail())

	if assert.NoError(t, err) {
		assert.Equal(t, account.ID, savedAccount.ID)
		assert.Equal(t, account.GetName(), savedAccount.GetName())
		assert.Equal(t, account.GetEmail(), savedAccount.GetEmail())
		assert.Equal(t, account.GetCPF(), savedAccount.GetCPF())
		assert.Equal(t, account.GetCarPlate(), savedAccount.GetCarPlate())
		assert.Equal(t, account.IsPassenger, savedAccount.IsPassenger)
		assert.Equal(t, account.IsDriver, savedAccount.IsDriver)
	}

	savedAccount, err = repo.GetAccountByID(account.ID)

	if assert.NoError(t, err) {
		assert.Equal(t, account.ID, savedAccount.ID)
		assert.Equal(t, account.GetName(), savedAccount.GetName())
		assert.Equal(t, account.GetEmail(), savedAccount.GetEmail())
		assert.Equal(t, account.GetCPF(), savedAccount.GetCPF())
		assert.Equal(t, account.GetCarPlate(), savedAccount.GetCarPlate())
		assert.Equal(t, account.IsPassenger, savedAccount.IsPassenger)
		assert.Equal(t, account.IsDriver, savedAccount.IsDriver)
	}
}

func Test_GetAccountByEmail_AndAccountIsNotFound(t *testing.T) {
	repo := di.NewAccountPostgresRepository()

	email := "invalid9999@mail.com"

	account, err := repo.GetAccountByEmail(email)

	assert.Nil(t, account)
	assert.Error(t, err)
}

func Test_GetAccountByID_AndAccountIsNotFound(t *testing.T) {
	repo := di.NewAccountPostgresRepository()

	accountId := uuid.New().String()

	account, err := repo.GetAccountByID(accountId)

	assert.Nil(t, account)
	assert.Error(t, err)
}

func getAccount() *entity.Account {
	account, err := entity.NewAccount(
		uuid.NewString(),
		"John Doe",
		fmt.Sprintf("john_%d@mail.com", rand.Int()),
		"97456321558",
		"",
		"secret123",
		true,
		false,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return account
}
