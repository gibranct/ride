package gateway_test

import (
	"testing"

	"github.com.br/gibranct/payment/internal/infra/gateway"
	"github.com/stretchr/testify/assert"
)

func Test_CreateTransaction_ReturnsError_WhenAmountIsZero(t *testing.T) {
	mockCielo := gateway.NewPaymentGatewayCielo()
	input := gateway.PaymentGatewayInput{
		Amount: 0,
	}

	output, err := mockCielo.CreateTransaction(input)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.Equal(t, "payment failed", err.Error())
}

func Test_CreateTransaction_ReturnsSuccess_WhenAmountIsAnOddNumber(t *testing.T) {
	mockCielo := gateway.NewPaymentGatewayCielo()
	input := gateway.PaymentGatewayInput{
		Amount: 9999999,
	}

	output, err := mockCielo.CreateTransaction(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, "1234567890", output.TID)
	assert.Equal(t, "123456", output.AuthorizationCode)
	assert.Equal(t, "approved", output.Status)
}
