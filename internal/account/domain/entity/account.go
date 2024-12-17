package domain

import (
	"github.com.br/gibranct/ride/internal/account/domain/vo"
	"github.com/google/uuid"
)

type Account struct {
	ID          string
	name        *vo.Name
	email       *vo.Email
	cpf         *vo.CPF
	carPlate    *vo.CarPlate
	IsPassenger bool
	IsDriver    bool
	password    *vo.Password
}

func (a *Account) GetName() string {
	return a.name.GetValue()
}

func (a *Account) GetEmail() string {
	return a.email.GetValue()
}

func (a *Account) GetCPF() string {
	return a.cpf.GetValue()
}

func (a *Account) GetCarPlate() string {
	if a.carPlate == nil {
		return ""
	}
	return a.carPlate.GetValue()
}

func (a *Account) GetPassword() string {
	return a.password.Value
}

func NewAccount(
	accountId, name, email, cpf, carPlate, password string, isPassenger, isDriver bool,
) (*Account, error) {
	newName, err := vo.NewName(name)
	if err != nil {
		return nil, err
	}
	newEmail, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}
	newCPF, err := vo.NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	newCarPlate, err := vo.NewCarPlate(carPlate)
	if isDriver && err != nil {
		return nil, err
	}
	if isPassenger {
		newCarPlate = nil
	}
	validPassword, err := vo.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:          accountId,
		name:        newName,
		email:       newEmail,
		cpf:         newCPF,
		carPlate:    newCarPlate,
		password:    validPassword,
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
