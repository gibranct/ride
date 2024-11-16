package main

type AccountDatabaseEntity struct {
	ID          string
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	IsPassenger bool
	IsDriver    bool
	Password    string
}

func (e *AccountDatabaseEntity) ToAccount() (*Account, error) {
	return NewAccount(
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
