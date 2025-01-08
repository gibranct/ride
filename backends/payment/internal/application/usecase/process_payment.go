package usecase

import (
	"context"

	"github.com.br/gibranct/payment/internal/domain"
	"github.com.br/gibranct/payment/internal/infra/fallback"
	"github.com.br/gibranct/payment/internal/infra/gateway"
	"github.com.br/gibranct/payment/internal/infra/repository"
)

type IProcessPayment interface {
	Execute(ctx context.Context, input ProcessPaymentInput) error
}

type ProcessPayment struct {
	paymentProcessor      fallback.PaymentProcessor
	transactionRepository repository.TransactionRepository
}

func (pp *ProcessPayment) Execute(ctx context.Context, input ProcessPaymentInput) error {
	inputTransaction := gateway.PaymentGatewayInput{
		CardHolder:       "Cliente Exemplo",
		CreditCardNumber: "4012001037141112",
		ExpDate:          "05/2027",
		CVV:              "123",
		Amount:           input.Amount,
	}
	transaction := domain.CreateTransaction(input.RideId, input.Amount)
	output, err := pp.paymentProcessor.ProcessPayment(inputTransaction)
	if err != nil {
		return err
	}
	if output.Status == "approved" {
		transaction.Pay()
		return pp.transactionRepository.SaveTransaction(ctx, *transaction)
	}
	return nil
}

type ProcessPaymentInput struct {
	RideId string  `json:"ride_id"`
	Amount float64 `json:"amount"`
}

func NewProcessPaymentUseCase(
	transactionRepository repository.TransactionRepository,
	paymentProcessor fallback.PaymentProcessor,
) *ProcessPayment {
	return &ProcessPayment{
		transactionRepository: transactionRepository,
		paymentProcessor:      paymentProcessor,
	}
}
