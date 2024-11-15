package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

const PORT = "127.0.0.1:3333"

type Account struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
	Password    string `json:"password"`
}

func main() {
	e := echo.New()

	e.POST("/sign-up", SignUp)

	e.Logger.Fatal(e.Start(PORT))
}

func SignUp(c echo.Context) error {
	saveQuery := "insert into gct.account (account_id, name, email, cpf, car_plate, is_passenger, is_driver, password) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	var input Account

	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	id := uuid.NewString()
	var result any

	account := Account{}
	conn.QueryRow(context.Background(), "select account_id, email from gct.account where email = $1", input.Email).Scan(
		&account.ID, &account.Email,
	)

	if account.ID != "" {
		response := map[string]any{"message": -4}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	matchCarPlate := regexp.MustCompile("[A-Z]{3}[0-9]{4}").MatchString(input.CarPlate)

	matchName := regexp.MustCompile("[a-zA-Z] [a-zA-Z]+").MatchString(input.Name)
	if !matchName {
		response := map[string]any{"message": -3}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	matchEmail := regexp.MustCompile("^(.+)@(.+)$").MatchString(input.Email)
	if !matchEmail {
		response := map[string]any{"message": -2}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if !validateCPF(input.CPF) {
		response := map[string]any{"message": -1}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if input.IsDriver && !matchCarPlate {
		response := map[string]any{"message": -5}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	args := []any{
		id, input.Name, input.Email, input.CPF, input.CarPlate, input.IsPassenger, input.IsDriver, input.Password,
	}
	conn.Exec(context.Background(), saveQuery, args...)
	obj := struct {
		AccountId string `json:"accountId"`
	}{
		AccountId: id,
	}
	result = obj

	return c.JSON(http.StatusCreated, result)
}
