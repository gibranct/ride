package usecase_test

import (
	"errors"
	"testing"

	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com.br/gibranct/payment/internal/domain"
	"github.com.br/gibranct/payment/internal/infra/fallback"
	"github.com.br/gibranct/payment/internal/infra/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentProcessor struct {
	mock.Mock
}

func (m *MockPaymentProcessor) ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*gateway.PaymentGatewayOutput), nil
	}
	return nil, args.Error(1)
}

func (m *MockPaymentProcessor) Next() fallback.PaymentProcessor {
	return nil
}

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) SaveTransaction(transaction domain.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetTransactionById(transactionId string) (*domain.Transaction, error) {
	args := m.Called(transactionId)
	if args.Get(0) == nil {
		return nil, nil
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func Test_Execute_ReturnsError_WhenPaymentProcessorFails(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	mockRepo := new(MockTransactionRepository)
	useCase := usecase.NewProcessPaymentUseCase(mockRepo, mockProcessor)

	input := usecase.ProcessPaymentInput{
		RideId: "ride123",
		Amount: 100.0,
	}

	mockProcessor.On("ProcessPayment", mock.Anything).Return(nil, errors.New("processing error"))

	err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, "processing error", err.Error())
	mockProcessor.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "SaveTransaction")
}

func Test_Execute_SavesTransaction_WhenPaymentIsApproved(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	mockRepo := new(MockTransactionRepository)
	useCase := usecase.NewProcessPaymentUseCase(mockRepo, mockProcessor)

	input := usecase.ProcessPaymentInput{
		RideId: "ride123",
		Amount: 100.0,
	}

	output := &gateway.PaymentGatewayOutput{
		Status: "approved",
	}

	mockProcessor.On("ProcessPayment", mock.Anything).Return(output, nil)
	mockRepo.On("SaveTransaction", mock.Anything).Return(nil)

	err := useCase.Execute(input)

	assert.Nil(t, err)
	mockProcessor.AssertExpectations(t)
	mockRepo.AssertCalled(t, "SaveTransaction", mock.Anything)
}

func Test_Execute_DoesNotSaveTransaction_WhenPaymentIsNotApproved(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	mockRepo := new(MockTransactionRepository)
	useCase := usecase.NewProcessPaymentUseCase(mockRepo, mockProcessor)

	input := usecase.ProcessPaymentInput{
		RideId: "ride123",
		Amount: 100.0,
	}

	output := &gateway.PaymentGatewayOutput{
		Status: "declined",
	}

	mockProcessor.On("ProcessPayment", mock.Anything).Return(output, nil)

	err := useCase.Execute(input)

	assert.Nil(t, err)
	mockProcessor.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "SaveTransaction", mock.Anything)
}
