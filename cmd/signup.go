package main

import (
	"context"
	"fmt"
	"os"
	"regexp"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SignUpOutput struct {
	AccountId string `json:"accountId"`
}

func GetAccountByEmail(email string) (*Account, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	account := Account{}
	conn.QueryRow(context.Background(), "select account_id, email from gct.account where email = $1", email).Scan(
		&account.ID, &account.Email,
	)

	return &account, nil
}

func GetAccountByID(id string) (*Account, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	var account Account
	conn.QueryRow(context.Background(), "select account_id, name, email, cpf, car_plate, is_passenger, is_driver from gct.account where account_id = $1", id).Scan(
		&account.ID, &account.Name, &account.Email, &account.CPF, &account.CarPlate, &account.IsPassenger, &account.IsDriver,
	)

	return &account, nil
}

func SaveAccount(account Account) error {
	saveQuery := "insert into gct.account (account_id, name, email, cpf, car_plate, is_passenger, is_driver, password) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	args := []any{
		account.ID, account.Name, account.Email, account.CPF, account.CarPlate, account.IsPassenger, account.IsDriver, account.Password,
	}
	_, err = conn.Exec(context.Background(), saveQuery, args...)

	return err
}

func SignUp(input Account) (*SignUpOutput, error) {
	input.ID = uuid.NewString()

	account, err := GetAccountByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if account.ID != "" {
		return nil, fmt.Errorf("duplicated account")
	}

	matchName := regexp.MustCompile("[a-zA-Z] [a-zA-Z]+").MatchString(input.Name)
	if !matchName {
		return nil, fmt.Errorf("invalid name")
	}

	matchEmail := regexp.MustCompile("^(.+)@(.+)$").MatchString(input.Email)
	if !matchEmail {
		return nil, fmt.Errorf("invalid email")
	}

	if !validateCPF(input.CPF) {
		return nil, fmt.Errorf("invalid cpf")
	}

	matchCarPlate := regexp.MustCompile("[A-Z]{3}[0-9]{4}").MatchString(input.CarPlate)
	if input.IsDriver && !matchCarPlate {
		return nil, fmt.Errorf("invalid car plate")
	}

	err = SaveAccount(input)

	if err != nil {
		return nil, err
	}

	return &SignUpOutput{
		AccountId: input.ID,
	}, nil
}
