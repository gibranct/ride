package test

import (
	"context"
	"log"
	"net/http/httptest"
	"os"
	"time"

	di "github.com.br/gibranct/payment/internal/infra/DI"
	myHttp "github.com.br/gibranct/payment/internal/infra/http"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"github.com/testcontainers/testcontainers-go/wait"
)

var scripts = []string{
	"../../create_payment.sql",
}

var testServer *httptest.Server

func setup() func() {
	dbContainer, connString := getDBConnStrAndContainer()
	rabbitContainer, rabbitURI := getRabbitContainer()
	err := os.Setenv("DATABASE_URL", connString)
	if err != nil {
		log.Fatalf("Error setting environment variable: %v", err)
	}

	err = os.Setenv("RABBITMQ_URI", rabbitURI)
	if err != nil {
		log.Fatalf("Error setting environment variable: %v", err)
	}

	httpServer := myHttp.NewHttpServer(di.NewProcessPayment())
	httpServer.SetUpRoutes()

	testServer = httptest.NewServer(httpServer.GetHandler())

	return func() {
		err := dbContainer.Terminate(context.Background())
		if err != nil {
			log.Printf("Error terminating database container: %v", err)
		}
		err = rabbitContainer.Terminate(context.Background())
		if err != nil {
			log.Printf("Error terminating rabbitmq container: %v", err)
		}
		testServer.Close()
	}
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

func getRabbitContainer() (*rabbitmq.RabbitMQContainer, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rabbitmqContainer, err := rabbitmq.Run(ctx,
		"rabbitmq:3.12.11-management-alpine",
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}
	connStr, err := rabbitmqContainer.AmqpURL(ctx)
	if err != nil {
		log.Fatalf("failed to get AMQP URL: %s", err)
	}

	return rabbitmqContainer, connStr
}
