package model

import (
	"time"

	"github.com.br/gibranct/payment/internal/domain"
)

type TransactionModel struct {
	TransactionID string  `db:"transaction_id"`
	RideId        string  `db:"ride_id"`
	Amount        float64 `db:"amount"`
	Status        string  `db:"status"`
	Date          string  `db:"date"`
}

func (t *TransactionModel) ToTransaction() *domain.Transaction {
	date, _ := time.Parse(time.RFC3339, t.Date)
	return domain.NewTransaction(
		t.TransactionID,
		t.RideId,
		t.Amount,
		t.Status,
		date,
	)
}
