package domain_test

import (
	"testing"
	"time"

	"github.com.br/gibranct/payment/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	transactionId := "123456"
	rideId := "789012"
	amount := 50.0
	status := "pending"
	date := time.Now()

	transaction := domain.NewTransaction(transactionId, rideId, amount, status, date)

	assert.NotNil(t, transaction)
	assert.Equal(t, transactionId, transaction.TransactionId)
	assert.Equal(t, rideId, transaction.RideId)
	assert.Equal(t, amount, transaction.GetAmount())
	assert.Equal(t, status, transaction.GetStatus())
	assert.Equal(t, date, transaction.GetDate())
}

func TestNewTransactionWithZeroAmount(t *testing.T) {
	transactionId := "123456"
	rideId := "789012"
	amount := 0.0
	status := "pending"
	date := time.Now()

	transaction := domain.NewTransaction(transactionId, rideId, amount, status, date)

	assert.NotNil(t, transaction)
	assert.Equal(t, transactionId, transaction.TransactionId)
	assert.Equal(t, rideId, transaction.RideId)
	assert.Equal(t, amount, transaction.GetAmount())
	assert.Equal(t, status, transaction.GetStatus())
	assert.Equal(t, date, transaction.GetDate())
}

func TestNewTransactionWithNegativeAmount(t *testing.T) {
	transactionId := "123456"
	rideId := "789012"
	amount := -50.0
	status := "pending"
	date := time.Now()

	transaction := domain.NewTransaction(transactionId, rideId, amount, status, date)

	assert.NotNil(t, transaction)
	assert.Equal(t, transactionId, transaction.TransactionId)
	assert.Equal(t, rideId, transaction.RideId)
	assert.Equal(t, amount, transaction.GetAmount())
	assert.Equal(t, status, transaction.GetStatus())
	assert.Equal(t, date, transaction.GetDate())
}

func TestCreateGeneratesValidUUID(t *testing.T) {
	rideId := "123456"
	amount := 50.0

	transaction := domain.CreateTransaction(rideId, amount)

	assert.NotNil(t, transaction)
	assert.NotEmpty(t, transaction.TransactionId)
	_, err := uuid.Parse(transaction.TransactionId)
	assert.NoError(t, err, "TransactionId should be a valid UUID")
	assert.Equal(t, rideId, transaction.RideId)
	assert.Equal(t, amount, transaction.GetAmount())
	assert.Equal(t, "waiting_payment", transaction.GetStatus())
	assert.WithinDuration(t, time.Now(), transaction.GetDate(), time.Second)
}

func TestCreateTransactionStatus(t *testing.T) {
	rideId := "test_ride_123"
	amount := 50.0

	transaction := domain.CreateTransaction(rideId, amount)

	assert.NotNil(t, transaction)
	assert.Equal(t, "waiting_payment", transaction.GetStatus())
}

func TestCreateSetsCurrentDate(t *testing.T) {
	rideId := "123456"
	amount := 50.0
	before := time.Now()
	transaction := domain.CreateTransaction(rideId, amount)
	after := time.Now()

	assert.NotNil(t, transaction)
	assert.Equal(t, rideId, transaction.RideId)
	assert.Equal(t, amount, transaction.GetAmount())
	assert.Equal(t, "waiting_payment", transaction.GetStatus())
	assert.True(t, transaction.GetDate().After(before) || transaction.GetDate().Equal(before))
	assert.True(t, transaction.GetDate().Before(after) || transaction.GetDate().Equal(after))
}

func TestChangeToPaidStatus(t *testing.T) {
	transactionId := "123456"
	rideId := "789012"
	amount := 50.0
	status := "pending"
	date := time.Now()

	transaction := domain.NewTransaction(transactionId, rideId, amount, status, date)

	assert.NotNil(t, transaction)
	assert.Equal(t, "pending", transaction.GetStatus())

	transaction.Pay()

	assert.Equal(t, "paid", transaction.GetStatus())
}
