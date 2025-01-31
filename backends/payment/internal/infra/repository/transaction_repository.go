package repository

import (
	"context"

	"github.com.br/gibranct/payment/internal/domain"
	"github.com.br/gibranct/payment/internal/infra/database"
	"github.com.br/gibranct/payment/internal/infra/repository/model"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, transaction domain.Transaction) error
	GetTransactionById(ctx context.Context, transactionId string) (*domain.Transaction, error)
}

type TransactionRepositoryDatabase struct {
	connection database.DatabaseConnection
}

func (repo TransactionRepositoryDatabase) SaveTransaction(ctx context.Context, transaction domain.Transaction) error {
	saveQuery := "insert into gct.transaction (transaction_id, ride_id, amount, status, date) values ($1, $2, $3, $4, $5)"
	args := []any{
		transaction.TransactionId, transaction.RideId, transaction.GetAmount(), transaction.GetStatus(), transaction.GetDate(),
	}
	return repo.connection.ExecContext(context.Background(), saveQuery, args...)
}

func (repo TransactionRepositoryDatabase) GetTransactionById(ctx context.Context, transactionId string) (*domain.Transaction, error) {
	var transactionModel model.TransactionModel
	query := "select transaction_id, ride_id, amount, status, date from gct.transaction where transaction_id = $1"
	err := repo.connection.QueryWithContext(context.Background(), &transactionModel, query, transactionId)
	if err != nil {
		return nil, err
	}
	return transactionModel.ToTransaction(), nil
}

func NewTransactionRepository(connection database.DatabaseConnection) TransactionRepository {
	return &TransactionRepositoryDatabase{connection: connection}
}
