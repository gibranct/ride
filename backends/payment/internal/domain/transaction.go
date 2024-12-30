package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionId string
	RideId        string
	amount        float64
	status        string
	date          *time.Time
}

func NewTransaction(transactionId, rideId string, amount float64, status string, date time.Time) *Transaction {
	return &Transaction{
		TransactionId: transactionId,
		RideId:        rideId,
		amount:        amount,
		status:        status,
		date:          &date,
	}
}

func CreateTransaction(rideId string, amount float64) *Transaction {
	transactionId := uuid.NewString()
	date := time.Now()
	status := "waiting_payment"
	return NewTransaction(transactionId, rideId, amount, status, date)
}

func (t *Transaction) Pay() {
	t.status = "paid"
}

func (t *Transaction) GetStatus() string {
	return t.status
}

func (t *Transaction) GetDate() time.Time {
	return *t.date
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}
