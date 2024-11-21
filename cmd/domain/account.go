package domain

import "github.com/google/uuid"

type Account struct {
	ID          string
	name        *Name
	email       *Email
	cpf         *CPF
	carPlate    *CarPlate
	IsPassenger bool
	IsDriver    bool
	Password    string
}

func (a *Account) GetName() string {
	return a.name.value
}

func (a *Account) GetEmail() string {
	return a.email.value
}

func (a *Account) GetCPF() string {
	return a.cpf.value
}

func (a *Account) GetCarPlate() string {
	if a.carPlate == nil {
		return ""
	}
	return a.carPlate.value
}

func NewAccount(
	accountId, name, email, cpf, carPlate, password string, isPassenger, isDriver bool,
) (*Account, error) {
	newName, err := NewName(name)
	if err != nil {
		return nil, err
	}
	newEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	newCPF, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	newCarPlate, err := NewCarPlate(carPlate)
	if isDriver && err != nil {
		return nil, err
	}
	if isPassenger {
		newCarPlate = nil
	}
	return &Account{
		ID:          accountId,
		name:        newName,
		email:       newEmail,
		cpf:         newCPF,
		carPlate:    newCarPlate,
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
