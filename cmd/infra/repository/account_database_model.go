package repository

import "github.com.br/gibranct/ride/cmd/domain"

type AccountDatabaseModel struct {
	ID          string
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	IsPassenger bool
	IsDriver    bool
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
