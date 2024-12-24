package fallback

import "github.com.br/gibranct/ride/internal/payment/infra/gateway"

type PjBankPaymentProcessor struct {
	paymentProcessor PaymentProcessor
}

func NewPjBankPaymentProcessor(next PaymentProcessor) *PjBankPaymentProcessor {
	return &PjBankPaymentProcessor{
		paymentProcessor: next,
	}
}

func (c *PjBankPaymentProcessor) Next() PaymentProcessor {
	return c.paymentProcessor
}

func (c *PjBankPaymentProcessor) ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error) {
	output, err := c.paymentProcessor.ProcessPayment(input)
	if err != nil && c.paymentProcessor.Next() == nil {
		return nil, err
	} else if err != nil {
		return c.paymentProcessor.Next().ProcessPayment(input)
	}
	return output, nil
}
