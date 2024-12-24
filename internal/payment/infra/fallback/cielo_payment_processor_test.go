package fallback_test

import (
	"errors"
	"testing"

	"github.com.br/gibranct/ride/internal/payment/infra/fallback"
	"github.com.br/gibranct/ride/internal/payment/infra/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentProcessor struct {
	mock.Mock
}

func (m *MockPaymentProcessor) ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gateway.PaymentGatewayOutput), args.Error(1)
}

func (m *MockPaymentProcessor) Next() fallback.PaymentProcessor {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(fallback.PaymentProcessor)
}

func Test_CieloProcessPayment_ReturnsError_WhenProcessorFailsAndNextIsNil(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{}
	expectedError := errors.New("processing error")

	mockProcessor.On("ProcessPayment", input).Return(nil, expectedError)
	mockProcessor.On("Next").Return(nil)

	cieloProcessor := fallback.NewCieloPaymentProcessor(mockProcessor)

	output, err := cieloProcessor.ProcessPayment(input)

	assert.Nil(t, output)
	assert.Equal(t, expectedError, err)
	mockProcessor.AssertExpectations(t)
}

func Test_CieloProcessPayment_ReturnsOutputFromNext_WhenProcessorFailsAndNextIsNotNil(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	mockNextProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{}
	expectedOutput := &gateway.PaymentGatewayOutput{}
	expectedError := errors.New("processing error")

	mockProcessor.On("ProcessPayment", input).Return(nil, expectedError)
	mockProcessor.On("Next").Return(mockNextProcessor)
	mockNextProcessor.On("ProcessPayment", input).Return(expectedOutput, nil)

	cieloProcessor := fallback.NewCieloPaymentProcessor(mockProcessor)

	output, err := cieloProcessor.ProcessPayment(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	mockProcessor.AssertExpectations(t)
	mockNextProcessor.AssertExpectations(t)
}

func Test_CieloProcessPayment_ReturnsOutputSuccessfully_WhenProcessorSucceeds(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)
	input := gateway.PaymentGatewayInput{}
	expectedOutput := &gateway.PaymentGatewayOutput{}

	mockProcessor.On("ProcessPayment", input).Return(expectedOutput, nil)

	cieloProcessor := fallback.NewCieloPaymentProcessor(mockProcessor)

	output, err := cieloProcessor.ProcessPayment(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	mockProcessor.AssertExpectations(t)
}

func Test_NewCieloPaymentProcessor_WithNilNextProcessor(t *testing.T) {
	cieloProcessor := fallback.NewCieloPaymentProcessor(nil)

	assert.NotNil(t, cieloProcessor)
	assert.Nil(t, cieloProcessor.Next())
}

func Test_NewCieloPaymentProcessor_CreatesProcessorWithValidNext(t *testing.T) {
	mockProcessor := new(MockPaymentProcessor)

	cieloProcessor := fallback.NewCieloPaymentProcessor(mockProcessor)

	assert.NotNil(t, cieloProcessor)
	assert.Equal(t, mockProcessor, cieloProcessor.Next())
}
