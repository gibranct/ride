package main

import "github.com/google/uuid"

type Account struct {
	ID          string `json:"id"`
	name        *Name
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
	Password    string `json:"password"`
}

func (a *Account) GetName() string {
	return a.name.value
}

func NewAccount(
	accountId, name, email, cpf, carPlate, password string, isPassenger, isDriver bool,
) (*Account, error) {
	newName, err := NewName(name)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:          accountId,
		name:        newName,
		Email:       email,
		CPF:         cpf,
		CarPlate:    carPlate,
		Password:    password,
		IsPassenger: isPassenger,
		IsDriver:    isDriver,
	}, nil
}

func CreateAccount(
	name, email, cpf, carPlate, password string, isPassenger, isDriver bool,
) (*Account, error) {
	accountId := uuid.NewString()
	return NewAccount(accountId, name, email, cpf, carPlate, password, isPassenger, isDriver)
}
