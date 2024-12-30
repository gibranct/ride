package fallback

import (
	"github.com.br/gibranct/payment/internal/infra/gateway"
)

type PaymentProcessor interface {
	Next() PaymentProcessor
	ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error)
}
