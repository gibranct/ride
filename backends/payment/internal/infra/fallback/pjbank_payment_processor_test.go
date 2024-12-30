package fallback_test

import (
	"errors"
	"testing"

	"github.com.br/gibranct/payment/internal/infra/fallback"
	"github.com.br/gibranct/payment/internal/infra/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_PJBankProcessPayment_ReturnsError_WhenCreateTransactionFails_AndNoNextProcessor(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockGateway.On("CreateTransaction", mock.Anything).Return(nil, errors.New("transaction error"))

	processor := fallback.NewPjBankPaymentProcessor(nil, mockGateway)
	input := gateway.PaymentGatewayInput{Amount: 100}

	output, err := processor.ProcessPayment(input)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.Equal(t, "transaction error", err.Error())
	mockGateway.AssertExpectations(t)
}

func Test_PJBankProcessPayment_CallsNextProcessor_WhenCreateTransactionFails_AndNextProcessorExists(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockGateway.On("CreateTransaction", mock.Anything).Return(nil, errors.New("transaction error"))

	mockNextProcessor := new(MockPaymentProcessor)
	mockNextProcessor.On("ProcessPayment", mock.Anything).Return(&gateway.PaymentGatewayOutput{TID: "next_tid"}, nil)

	processor := fallback.NewPjBankPaymentProcessor(mockNextProcessor, mockGateway)
	input := gateway.PaymentGatewayInput{Amount: 100}

	output, err := processor.ProcessPayment(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, "next_tid", output.TID)
	mockGateway.AssertExpectations(t)
	mockNextProcessor.AssertExpectations(t)
}

func Test_PJBankProcessPayment_ReturnsOutput_WhenCreateTransactionSucceeds(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockGateway.On("CreateTransaction", mock.Anything).Return(&gateway.PaymentGatewayOutput{TID: "success_tid"}, nil)

	processor := fallback.NewPjBankPaymentProcessor(nil, mockGateway)
	input := gateway.PaymentGatewayInput{Amount: 100}

	output, err := processor.ProcessPayment(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, "success_tid", output.TID)
	mockGateway.AssertExpectations(t)
}

func Test_PJBankProcessPayment_PropagatesError_WhenNoNextProcessorAvailable(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockGateway.On("CreateTransaction", mock.Anything).Return(nil, errors.New("create transaction error"))

	processor := fallback.NewPjBankPaymentProcessor(nil, mockGateway)
	input := gateway.PaymentGatewayInput{Amount: 100}

	output, err := processor.ProcessPayment(input)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.Equal(t, "create transaction error", err.Error())
	mockGateway.AssertExpectations(t)
}

func Test_PJBankProcessPayment_ReturnsOutputFromNextProcessor_WhenCreateTransactionFails_AndNextProcessorExists(t *testing.T) {
	mockGateway := new(MockPaymentGateway)
	mockGateway.On("CreateTransaction", mock.Anything).Return(nil, errors.New("transaction error"))

	mockNextProcessor := new(MockPaymentProcessor)
	mockNextProcessor.On("ProcessPayment", mock.Anything).Return(&gateway.PaymentGatewayOutput{TID: "next_processor_tid"}, nil)

	processor := fallback.NewPjBankPaymentProcessor(mockNextProcessor, mockGateway)
	input := gateway.PaymentGatewayInput{Amount: 100}

	output, err := processor.ProcessPayment(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, "next_processor_tid", output.TID)
	mockGateway.AssertExpectations(t)
	mockNextProcessor.AssertExpectations(t)
}
