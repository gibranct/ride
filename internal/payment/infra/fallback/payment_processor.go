package fallback

import (
	"github.com.br/gibranct/ride/internal/payment/infra/gateway"
)

type PaymentProcessor interface {
	Next() PaymentProcessor
	ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error)
}
