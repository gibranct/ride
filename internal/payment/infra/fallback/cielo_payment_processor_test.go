package fallback_test

import (
	"errors"
	"testing"

	"github.com.br/gibranct/ride/internal/payment/infra/fallback"
	"github.com.br/gibranct/ride/internal/payment/infra/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentGateway struct {
	mock.Mock
}

func (m *MockPaymentGateway) CreateTransaction(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*gateway.PaymentGatewayOutput), nil
	}
	return nil, args.Error(1)
}

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
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(fallback.PaymentProcessor)
}

func Test_CieloProcessPayment_ReturnsError_WhenCreateTransactionFailsAndNoNextProcessor(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	input := gateway.PaymentGatewayInput{
		Amount: 100,
	}
	expectedError := errors.New("transaction error")
	mockGateway.On("CreateTransaction", input).Return(nil, expectedError)

	processor := fallback.NewCieloPaymentProcessor(nil, mockGateway)

	output, err := processor.ProcessPayment(input)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockGateway.AssertExpectations(t)
}

func Test_CieloProcessPayment_ProceedsToNextProcessor_WhenCreateTransactionFailsAndNextProcessorExists(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockNextProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{
		Amount: 100,
	}
	expectedError := errors.New("transaction error")
	mockGateway.On("CreateTransaction", input).Return(nil, expectedError)
	mockNextProcessor.On("ProcessPayment", input).Return(&gateway.PaymentGatewayOutput{TID: "next_tid"}, nil)

	processor := fallback.NewCieloPaymentProcessor(mockNextProcessor, mockGateway)

	output, err := processor.ProcessPayment(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, "next_tid", output.TID)
	mockGateway.AssertExpectations(t)
	mockNextProcessor.AssertExpectations(t)
}

func Test_CieloProcessPayment_ReturnsOutput_WhenCreateTransactionSucceeds(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	input := gateway.PaymentGatewayInput{
		Amount: 100,
	}
	expectedOutput := &gateway.PaymentGatewayOutput{TID: "successful_tid"}
	mockGateway.On("CreateTransaction", input).Return(expectedOutput, nil)

	processor := fallback.NewCieloPaymentProcessor(nil, mockGateway)

	output, err := processor.ProcessPayment(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	mockGateway.AssertExpectations(t)
}

func Test_CieloProcessPayment_ChainsToMultipleProcessors_WhenEachReturnsError(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockNextProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{
		Amount: 100,
	}
	expectedError := errors.New("transaction error")
	mockGateway.On("CreateTransaction", input).Return(nil, expectedError)
	mockNextProcessor.On("ProcessPayment", input).Return(&gateway.PaymentGatewayOutput{TID: "final_tid"}, nil)

	processor1 := fallback.NewCieloPaymentProcessor(mockNextProcessor, mockGateway)
	processor := fallback.NewCieloPaymentProcessor(processor1, mockGateway)

	output, err := processor.ProcessPayment(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, "final_tid", output.TID)
	mockGateway.AssertExpectations(t)
	mockNextProcessor.AssertExpectations(t)
}

func Test_CieloProcessPayment_DoesNotCallNextProcessor_WhenCreateTransactionSucceeds(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockNextProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{
		Amount: 100,
	}
	expectedOutput := &gateway.PaymentGatewayOutput{TID: "successful_tid"}
	mockGateway.On("CreateTransaction", input).Return(expectedOutput, nil)

	processor := fallback.NewCieloPaymentProcessor(mockNextProcessor, mockGateway)

	output, err := processor.ProcessPayment(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	mockGateway.AssertExpectations(t)
	mockNextProcessor.AssertNotCalled(t, "ProcessPayment", input)
}
