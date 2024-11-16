package main

type GetAccount struct {
	accountDAO AccountDAO
}

type GetAccountOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
}

func (gc *GetAccount) Execute(accountId string) (*GetAccountOutput, error) {
	account, err := gc.accountDAO.GetAccountByID(accountId)

	if err != nil {
		return nil, err
	}

	return &GetAccountOutput{
		ID:          accountId,
		Name:        account.GetName(),
		Email:       account.Email,
		CPF:         account.CPF,
		CarPlate:    account.CarPlate,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil
}

func NewGetAccountCase(accountDAO AccountDAO) *GetAccount {
	return &GetAccount{
		accountDAO: accountDAO,
	}
}
