package main

type GetAccount struct {
	accountDAO AccountRepository
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
		Email:       account.GetEmail(),
		CPF:         account.GetCPF(),
		CarPlate:    account.GetCarPlate(),
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil
}

func NewGetAccountCase(accountDAO AccountRepository) *GetAccount {
	return &GetAccount{
		accountDAO: accountDAO,
	}
}