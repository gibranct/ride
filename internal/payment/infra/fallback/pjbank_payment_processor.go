package fallback

import "github.com.br/gibranct/ride/internal/payment/infra/gateway"

type PjBankPaymentProcessor struct {
	paymentProcessor PaymentProcessor
	pJBankGateway    gateway.PaymentGateway
}

func NewPjBankPaymentProcessor(next PaymentProcessor, pjBankGateway gateway.PaymentGateway) *PjBankPaymentProcessor {
	return &PjBankPaymentProcessor{
		paymentProcessor: next,
		pJBankGateway:    pjBankGateway,
	}
}

func (c *PjBankPaymentProcessor) Next() PaymentProcessor {
	return c.paymentProcessor
}

func (c *PjBankPaymentProcessor) ProcessPayment(input gateway.PaymentGatewayInput) (*gateway.PaymentGatewayOutput, error) {
	output, err := c.pJBankGateway.CreateTransaction(input)
	if err != nil && c.Next() == nil {
		return nil, err
	} else if err != nil {
		return c.Next().ProcessPayment(input)
	}
	return output, nil
}
