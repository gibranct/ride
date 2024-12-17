package application

import (
	"github.com.br/gibranct/ride/internal/payment/application/usecase"
	di "github.com.br/gibranct/ride/internal/payment/infra/DI"
)

type PaymentService struct {
	*usecase.ProcessPayment
}

type Application struct {
	*PaymentService
}

func NewApplication() *Application {
	return &Application{
		PaymentService: &PaymentService{
			ProcessPayment: di.NewProcessPayment(),
		},
	}
}
