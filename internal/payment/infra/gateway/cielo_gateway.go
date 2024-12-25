package gateway

import (
	"fmt"
)

type PaymentGatewayCielo struct{}

func (m *PaymentGatewayCielo) CreateTransaction(input PaymentGatewayInput) (*PaymentGatewayOutput, error) {
	fmt.Printf("processing payment cielo: %+v\n", input)
	if int(input.Amount)%2 == 0 {
		return nil, fmt.Errorf("payment failed")
	}
	return &PaymentGatewayOutput{
		TID:               "1234567890",
		AuthorizationCode: "123456",
		Status:            "approved",
	}, nil
}

func NewPaymentGatewayCielo() PaymentGateway {
	return &PaymentGatewayCielo{}
}
