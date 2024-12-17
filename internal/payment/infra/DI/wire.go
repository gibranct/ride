//go:build wireinject
// +build wireinject

package di

import (
	"github.com.br/gibranct/ride/internal/payment/application/usecase"
	"github.com/google/wire"
)

func NewProcessPayment() *usecase.ProcessPayment {
	wire.Build(
		usecase.NewProcessPaymentUseCase,
	)
	return &usecase.ProcessPayment{}
}
