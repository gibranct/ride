package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com.br/gibranct/payment/internal/domain"
	di "github.com.br/gibranct/payment/internal/infra/DI"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	teardown := setup()

	defer teardown()

	os.Exit(m.Run())
}

func TestSaveTransaction(t *testing.T) {
	repo := di.NewTransactionPostgresRepository()

	tid := uuid.NewString()
	rideId := uuid.NewString()
	transaction := domain.NewTransaction(tid, rideId, 100, "pending", time.Now().UTC())
	ctxBackground := context.Background()

	err := repo.SaveTransaction(ctxBackground, *transaction)
	if err != nil {
		t.Errorf("failed to save transaction: %s", err)
	}

	savedTransaction, err := repo.GetTransactionById(ctxBackground, tid)
	if err != nil {
		t.Errorf("failed to get saved transaction: %s", err)
	}

	if assert.NoError(t, err) {
		assert.Equal(t, transaction.TransactionId, savedTransaction.TransactionId)
		assert.Equal(t, transaction.RideId, savedTransaction.RideId)
		assert.Equal(t, transaction.GetAmount(), savedTransaction.GetAmount())
		assert.Equal(t, transaction.GetStatus(), savedTransaction.GetStatus())
		assert.Equal(t, transaction.GetDate(), savedTransaction.GetDate())
	}
}
