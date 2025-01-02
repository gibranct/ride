package test

import (
	"context"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var scripts = []string{
	"../../create_account.sql",
}

func getDBConnStrAndContainer() (*postgres.PostgresContainer, string) {
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
		log.Fatalf("failed to get connection string: %s", err)
	}

	return postgresC, connString
}
