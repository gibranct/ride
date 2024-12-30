package fallback

import (
	"github.com.br/gibranct/payment/internal/infra/gateway"
)

type CieloPaymentProcessor struct {
	paymentProcessor PaymentProcessor
	cieloGateway     gateway.PaymentGateway
}

func NewCieloPaymentProcessor(next PaymentProcessor, cieloGateway gateway.PaymentGateway) *CieloPaymentProcessor {
	return &CieloPaymentProcessor{
		paymentProcessor: next,
		cieloGateway:     cieloGateway,
	}
}

func (c *CieloPaymentProcessor) Next() PaymentProcessor {
	return c.paymentProcessor
}

func (c *CieloPaymentProcessor) ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error) {
	output, err := c.cieloGateway.CreateTransaction(input)
	if err != nil && c.Next() == nil {
		return nil, err
	} else if err != nil {
		return c.Next().ProcessPayment(input)
	}
	return output, nil
}
