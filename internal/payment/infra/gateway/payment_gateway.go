package gateway

type PaymentGateway interface {
	CreateTransaction(input PaymentGatewayInput) (*PaymentGatewayOutput, error)
}

type PaymentGatewayInput struct {
	CardHolder       string
	CreditCardNumber string
	ExpDate          string
	CVV              string
	Amount           float64
}

type PaymentGatewayOutput struct {
	TID               string
	AuthorizationCode string
	Status            string
}
