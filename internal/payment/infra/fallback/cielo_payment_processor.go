package fallback

import "github.com.br/gibranct/ride/internal/payment/infra/gateway"

type CieloPaymentProcessor struct {
	paymentProcessor PaymentProcessor
}

func NewCieloPaymentProcessor(next PaymentProcessor) *CieloPaymentProcessor {
	return &CieloPaymentProcessor{
		paymentProcessor: next,
	}
}

func (c *CieloPaymentProcessor) Next() PaymentProcessor {
	return c.paymentProcessor
}

func (c *CieloPaymentProcessor) ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error) {
	output, err := c.paymentProcessor.ProcessPayment(input)
	if err != nil && c.paymentProcessor.Next() == nil {
		return nil, err
	} else if err != nil {
		return c.paymentProcessor.Next().ProcessPayment(input)
	}
	return output, nil
}
