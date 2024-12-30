package repository_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com.br/gibranct/payment/internal/domain"
	"github.com.br/gibranct/payment/internal/infra/database"
	"github.com.br/gibranct/payment/internal/infra/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var scripts = []string{
	"../../../../../create_payment.sql",
}

func TestSaveTransaction(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a PostgreSQL container
	postgresC, err := postgres.Run(
		ctx,
		"postgres:15.7-alpine",
		postgres.WithInitScripts(scripts...),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		log.Fatalf("Could not start container: %s", err)
	}

	connString, err := postgresC.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %s", err)
	}
	defer postgresC.Terminate(ctx)

	// Set up the database connection
	dbConn := database.NewPostgresAdapter(connString)

	repo := repository.NewTransactionRepository(dbConn)

	// Create a test transaction
	tid := uuid.NewString()
	rideId := uuid.NewString()
	transaction := domain.NewTransaction(tid, rideId, 100, "pending", time.Now().UTC())

	// Test SaveTransaction
	err = repo.SaveTransaction(*transaction)
	if err != nil {
		t.Errorf("failed to save transaction: %s", err)
	}

	savedTransaction, err := repo.GetTransactionById(tid)
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
