package gateway

import (
	"fmt"
	"math/rand"
)

type PaymentGatewayPJBank struct{}

func (m *PaymentGatewayPJBank) CreateTransaction(input PaymentGatewayInput) (*PaymentGatewayOutput, error) {
	fmt.Printf("processing payment PJBank: %+v\n", input)
	if rand.Int()%2 == 0 {
		return nil, fmt.Errorf("payment failed")
	}
	return &PaymentGatewayOutput{
		TID:               "1234567890",
		AuthorizationCode: "123456",
		Status:            "approved",
	}, nil
}

func NewPaymentGatewayPJBank() *PaymentGatewayPJBank {
	return &PaymentGatewayPJBank{}
}
