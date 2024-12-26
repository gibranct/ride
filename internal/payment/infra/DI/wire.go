//go:build wireinject
// +build wireinject

package di

import (
	"github.com.br/gibranct/ride/internal/payment/application/usecase"
	"github.com.br/gibranct/ride/internal/payment/infra/database"
	"github.com.br/gibranct/ride/internal/payment/infra/fallback"
	"github.com.br/gibranct/ride/internal/payment/infra/gateway"
	"github.com.br/gibranct/ride/internal/payment/infra/repository"
	"github.com/google/wire"
)

var databaseSet = wire.NewSet(
	database.NewPostgresAdapter,
	wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)),
)

func NewProcessPayment() *usecase.ProcessPayment {
	wire.Build(
		NewPaymentProcessor,
		databaseSet,
		repository.NewTransactionRepository,
		usecase.NewProcessPaymentUseCase,
	)
	return &usecase.ProcessPayment{}
}

func NewPaymentProcessor() fallback.PaymentProcessor {
	pjBankProcessor := fallback.NewPjBankPaymentProcessor(nil, gateway.NewPaymentGatewayPJBank())
	cieloProcessor := fallback.NewCieloPaymentProcessor(pjBankProcessor, gateway.NewPaymentGatewayCielo())
	return cieloProcessor
}
