package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type AccountRepository interface {
	GetAccountByEmail(email string) (*Account, error)
	GetAccountByID(id string) (*Account, error)
	SaveAccount(account Account) error
}

type AccountRepositoryDatabase struct{}
type AccountRepositoryMemory struct {
	accounts []Account
}

func (dao AccountRepositoryDatabase) GetAccountByEmail(email string) (*Account, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	account := AccountDatabaseEntity{}
	conn.QueryRow(context.Background(), "select account_id, name, email, cpf, car_plate, is_passenger, is_driver from gct.account where email = $1", email).Scan(
		&account.ID, &account.Name, &account.Email, &account.CPF, &account.CarPlate, &account.IsPassenger, &account.IsDriver,
	)

	if account.ID == "" {
		return &Account{}, nil
	}

	return account.ToAccount()
}

func (dao AccountRepositoryDatabase) GetAccountByID(id string) (*Account, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	var account AccountDatabaseEntity
	conn.QueryRow(context.Background(), "select account_id, name, email, cpf, car_plate, is_passenger, is_driver from gct.account where account_id = $1", id).Scan(
		&account.ID, &account.Name, &account.Email, &account.CPF, &account.CarPlate, &account.IsPassenger, &account.IsDriver,
	)

	return account.ToAccount()
}

func (dao AccountRepositoryDatabase) SaveAccount(account Account) error {
	saveQuery := "insert into gct.account (account_id, name, email, cpf, car_plate, is_passenger, is_driver, password) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	args := []any{
		account.ID, account.GetName(), account.GetEmail(), account.GetCPF(), account.GetCarPlate(), account.IsPassenger, account.IsDriver, account.Password,
	}
	_, err = conn.Exec(context.Background(), saveQuery, args...)

	return err
}

func NewAccountDAO() AccountRepository {
	return &AccountRepositoryDatabase{}
}

func (dao *AccountRepositoryMemory) GetAccountByEmail(email string) (*Account, error) {
	for i := range dao.accounts {
		if dao.accounts[i].GetEmail() == email {
			return &dao.accounts[i], nil
		}
	}

	return &Account{}, nil
}

func (dao *AccountRepositoryMemory) GetAccountByID(id string) (*Account, error) {
	for i := range dao.accounts {
		if dao.accounts[i].ID == id {
			return &dao.accounts[i], nil
		}
	}

	return &Account{}, nil
}

func (dao *AccountRepositoryMemory) SaveAccount(account Account) error {
	dao.accounts = append(dao.accounts, account)
	return nil
}

func NewAccountRepositoryMemory() AccountRepository {
	return &AccountRepositoryMemory{}
}
