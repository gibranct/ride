package main

import (
	"fmt"
	"regexp"

	uuid "github.com/google/uuid"
)

type SignUpOutput struct {
	AccountId string `json:"accountId"`
}

func SignUp(input Account) (*SignUpOutput, error) {
	accountDAO := NewAccountDAO()
	input.ID = uuid.NewString()

	account, err := accountDAO.GetAccountByEmail(input.Email)

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

	err = accountDAO.SaveAccount(input)

	if err != nil {
		return nil, err
	}

	return &SignUpOutput{
		AccountId: input.ID,
	}, nil
}
