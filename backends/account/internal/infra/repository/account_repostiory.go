package repository

import (
	"context"

	domain "github.com.br/gibranct/account/internal/domain/entity"
	"github.com.br/gibranct/account/internal/infra/database"
	"github.com.br/gibranct/account/internal/infra/repository/model"
)

type AccountRepository interface {
	GetAccountByEmail(email string) (*domain.Account, error)
	GetAccountByID(id string) (*domain.Account, error)
	SaveAccount(account domain.Account) error
}

type AccountRepositoryDatabase struct {
	connection database.DatabaseConnection
}

func (dao AccountRepositoryDatabase) GetAccountByEmail(email string) (*domain.Account, error) {
	account := &model.AccountDatabaseModel{}
	query := "select account_id, password, name, email, cpf, car_plate, is_passenger, is_driver from gct.account where email = $1"
	err := dao.connection.QueryWithContext(context.Background(), account, query, email)
	if err != nil {
		return nil, err
	}
	return account.ToAccount()
}

func (dao AccountRepositoryDatabase) GetAccountByID(id string) (*domain.Account, error) {
	var account model.AccountDatabaseModel
	query := "select account_id, password, name, email, cpf, car_plate, is_passenger, is_driver from gct.account where account_id = $1"
	err := dao.connection.QueryWithContext(context.Background(), &account, query, id)
	if err != nil {
		return nil, err
	}

	return account.ToAccount()
}

func (dao AccountRepositoryDatabase) SaveAccount(account domain.Account) error {
	saveQuery := "insert into gct.account (account_id, name, email, cpf, car_plate, is_passenger, is_driver, password) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	args := []any{
		account.ID, account.GetName(), account.GetEmail(), account.GetCPF(), account.GetCarPlate(), account.IsPassenger, account.IsDriver, account.GetPassword(),
	}
	return dao.connection.ExecContext(context.Background(), saveQuery, args...)
}

func NewAccountRepository(conn database.DatabaseConnection) *AccountRepositoryDatabase {
	return &AccountRepositoryDatabase{
		connection: conn,
	}
}
