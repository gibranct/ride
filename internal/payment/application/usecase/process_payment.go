package usecase

import "fmt"

type ProcessPayment struct{}

func (pp *ProcessPayment) Execute(input ProcessPaymentInput) error {
	fmt.Printf("Processing payment: amount=%f, rideId=%s\n", input.Amount, input.RideId)
	return nil
}

type ProcessPaymentInput struct {
	RideId string
	Amount float64
}

func NewProcessPaymentUseCase() *ProcessPayment {
	return &ProcessPayment{}
}
