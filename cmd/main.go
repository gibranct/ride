package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sign-up", SignUp)

	server := http.Server{
		Addr:         PORT,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Handler:      mux,
	}

	fmt.Printf("Server listening on port %s...\n", PORT)

	err := server.ListenAndServe()

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func SignUp(w http.ResponseWriter, req *http.Request) {
	saveQuery := "insert into gct.account (account_id, name, email, cpf, car_plate, is_passenger, is_driver, password) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	var input Account

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
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

	if account.ID == "" {
		matchName := regexp.MustCompile("[a-zA-Z] [a-zA-Z]+").MatchString(input.Name)
		matchEmail := regexp.MustCompile("^(.+)@(.+)$").MatchString(input.Email)
		matchCarPlate := regexp.MustCompile("[A-Z]{3}[0-9]{4}").MatchString(input.CarPlate)
		if matchName {
			if matchEmail {
				if validateCPF(input.CPF) {
					if input.IsDriver {
						if matchCarPlate {
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
						} else {
							result = -5
						}
					} else {
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
					}
				} else {
					// invalid cpf
					result = -1
				}
			} else {
				// invalid email
				result = -2
			}

		} else {
			// invalid name
			result = -3
		}
	} else {
		// already exists
		result = -4
	}

	w.Header().Set("Content-Type", "application/json")

	if _, ok := result.(int); ok {
		w.WriteHeader(http.StatusUnprocessableEntity)

		err = json.NewEncoder(w).Encode(
			map[string]any{
				"message": result,
			},
		)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}
}
