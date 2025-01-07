package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com.br/gibranct/account/internal/application/usecase"
	"github.com.br/gibranct/account/internal/domain/entity"
	"github.com.br/gibranct/account/internal/domain/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAccountExecute_AccountNotFound(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	getAccountUseCase := usecase.NewGetAccountUseCase(mockRepo)
	ctxBackground := context.Background()

	accountId := "non-existent-id"
	mockRepo.On("GetAccountByID", ctxBackground, accountId).Return(nil, errors.ErrAccountNotFound)

	output, err := getAccountUseCase.Execute(ctxBackground, accountId)

	assert.Nil(t, output)
	assert.Equal(t, errors.ErrAccountNotFound, err)
	mockRepo.AssertCalled(t, "GetAccountByID", ctxBackground, accountId)
}

func Test_GetAccountExecute_ValidAccountId(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	getAccountUseCase := usecase.NewGetAccountUseCase(mockRepo)
	ctxBackground := context.Background()

	accountId := uuid.New().String()
	expectedAccount, err := entity.NewAccount(
		accountId, "Jane Doe", "jane@doe.com",
		"12345678909", "XYZ1234", "secret123", true, false,
	)
	assert.NoError(t, err)

	mockRepo.On("GetAccountByID", ctxBackground, accountId).Return(expectedAccount, nil)

	output, err := getAccountUseCase.Execute(ctxBackground, accountId)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, expectedAccount.ID, output.ID)
	assert.Equal(t, expectedAccount.GetName(), output.Name)
	assert.Equal(t, expectedAccount.GetEmail(), output.Email)
	assert.Equal(t, expectedAccount.GetCPF(), output.CPF)
	assert.Equal(t, expectedAccount.GetCarPlate(), output.CarPlate)
	assert.Equal(t, expectedAccount.IsPassenger, output.IsPassenger)
	assert.Equal(t, expectedAccount.IsDriver, output.IsDriver)
	mockRepo.AssertCalled(t, "GetAccountByID", ctxBackground, accountId)
}

func Test_GetAccountExecute_WithTimeout(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	getAccountUseCase := usecase.NewGetAccountUseCase(mockRepo)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	accountId := uuid.New().String()
	expectedAccount, err := entity.NewAccount(
		accountId, "Jane Doe", "jane@doe.com",
		"12345678909", "XYZ1234", "secret123", true, false,
	)
	assert.NoError(t, err)

	mockRepo.On("GetAccountByID", accountId).Return(expectedAccount, nil)

	select {
	case <-ctxTimeout.Done():
	default:
		getAccountUseCase.Execute(ctxTimeout, accountId)
		assert.Fail(t, "Context did not timeout as expected")
	}

}
