// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com.br/gibranct/payment/internal/infra/database"
	"github.com.br/gibranct/payment/internal/infra/fallback"
	"github.com.br/gibranct/payment/internal/infra/gateway"
	"github.com.br/gibranct/payment/internal/infra/repository"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewProcessPayment() *usecase.ProcessPayment {
	postgresAdapter := database.NewPostgresAdapter()
	transactionRepository := repository.NewTransactionRepository(postgresAdapter)
	paymentProcessor := NewPaymentProcessor()
	processPayment := usecase.NewProcessPaymentUseCase(transactionRepository, paymentProcessor)
	return processPayment
}

func NewTransactionPostgresRepository() repository.TransactionRepository {
	postgresAdapter := database.NewPostgresAdapter()
	transactionRepository := repository.NewTransactionRepository(postgresAdapter)
	return transactionRepository
}

// wire.go:

var databaseSet = wire.NewSet(database.NewPostgresAdapter, wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)))

func NewPaymentProcessor() fallback.PaymentProcessor {
	pjBankProcessor := fallback.NewPjBankPaymentProcessor(nil, gateway.NewPaymentGatewayPJBank())
	cieloProcessor := fallback.NewCieloPaymentProcessor(pjBankProcessor, gateway.NewPaymentGatewayCielo())
	return cieloProcessor
}
