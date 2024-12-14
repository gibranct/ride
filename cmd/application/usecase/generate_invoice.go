package usecase

import "fmt"

type GenerateInvoice struct{}

func (pp *GenerateInvoice) Execute(input GenerateInvoiceInput) error {
	fmt.Printf("Processing payment: amount=%f, rideId=%s\n", input.Amount, input.RideId)
	return nil
}

type GenerateInvoiceInput struct {
	RideId string
	Amount float64
}

func NewGenerateInvoiceUseCase() *GenerateInvoice {
	return &GenerateInvoice{}
}
