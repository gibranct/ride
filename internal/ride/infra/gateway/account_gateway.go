package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AccountGateway struct{}

type Account struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
}

type SignUpInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
	Password    string `json:"password"`
}

type SignUpOutput struct {
	AccountId string `json:"accountId"`
}

func (m *AccountGateway) GetAccount(accountId string) (*Account, error) {

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://127.0.0.1:3001/v1/accounts/%s", accountId), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	client.Timeout = 5 * time.Second

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching account: %s", response.Status)
	}

	var account Account
	err = json.NewDecoder(response.Body).Decode(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (m *AccountGateway) SignUp(input SignUpInput) (string, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:3001/sign-up", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Timeout = 5 * time.Second

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("error signing up account: %s", response.Status)
	}

	var output SignUpOutput
	err = json.NewDecoder(response.Body).Decode(&output)

	if err != nil {
		return "", err
	}

	return output.AccountId, nil
}

func NewAccountGateway() AccountGateway {
	return AccountGateway{}
}
