package model

import (
	"github.com.br/gibranct/ride/cmd/domain"
)

type AccountDatabaseModel struct {
	ID          string `db:"account_id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	CPF         string `db:"cpf"`
	CarPlate    string `db:"car_plate"`
	IsPassenger bool   `db:"is_passenger"`
	IsDriver    bool   `db:"is_driver"`
	Password    string
}

func (e *AccountDatabaseModel) ToAccount() (*domain.Account, error) {
	return domain.NewAccount(
		e.ID,
		e.Name,
		e.Email,
		e.CPF,
		e.CarPlate,
		e.Password,
		e.IsPassenger,
		e.IsDriver,
	)
}
