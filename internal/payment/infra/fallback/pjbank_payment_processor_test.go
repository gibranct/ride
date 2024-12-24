package fallback_test

import (
	"errors"
	"testing"

	"github.com.br/gibranct/ride/internal/payment/infra/fallback"
	"github.com.br/gibranct/ride/internal/payment/infra/gateway"
	"github.com/stretchr/testify/assert"
)

func Test_PjBankProcessPayment_ReturnsError_WhenProcessorFailsAndNextIsNil(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{}
	expectedError := errors.New("processing error")

	mockProcessor.On("ProcessPayment", input).Return(nil, expectedError)
	mockProcessor.On("Next").Return(nil)

	pjBankProcessor := fallback.NewPjBankPaymentProcessor(mockProcessor)

	output, err := pjBankProcessor.ProcessPayment(input)

	assert.Nil(t, output)
	assert.Equal(t, expectedError, err)
	mockProcessor.AssertExpectations(t)
}

func Test_PjBankProcessPayment_ReturnsOutputFromNext_WhenProcessorFailsAndNextIsNotNil(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	mockNextProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{}
	expectedOutput := &gateway.PaymentGatewayOutput{}
	expectedError := errors.New("processing error")

	mockProcessor.On("ProcessPayment", input).Return(nil, expectedError)
	mockProcessor.On("Next").Return(mockNextProcessor)
	mockNextProcessor.On("ProcessPayment", input).Return(expectedOutput, nil)

	pjBankProcessor := fallback.NewPjBankPaymentProcessor(mockProcessor)

	output, err := pjBankProcessor.ProcessPayment(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	mockProcessor.AssertExpectations(t)
	mockNextProcessor.AssertExpectations(t)
}

func Test_PjBankProcessPayment_ReturnsOutputSuccessfully_WhenProcessorSucceeds(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{}
	expectedOutput := &gateway.PaymentGatewayOutput{}

	mockProcessor.On("ProcessPayment", input).Return(expectedOutput, nil)

	pjBankProcessor := fallback.NewPjBankPaymentProcessor(mockProcessor)

	output, err := pjBankProcessor.ProcessPayment(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	mockProcessor.AssertExpectations(t)
}

func Test_NewPjBankPaymentProcessor_WithNilNextProcessor(t *testing.T) {
	pjBankProcessor := fallback.NewPjBankPaymentProcessor(nil)

	assert.NotNil(t, pjBankProcessor)
	assert.Nil(t, pjBankProcessor.Next())
}

func Test_NewPjBankPaymentProcessor_CreatesProcessorWithValidNext(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)

	pjBankProcessor := fallback.NewPjBankPaymentProcessor(mockProcessor)

	assert.NotNil(t, pjBankProcessor)
	assert.Equal(t, mockProcessor, pjBankProcessor.Next())
}
