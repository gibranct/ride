package main

type GetAccount struct {
	accountDAO AccountDAO
}

type Account struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
	Password    string `json:"password"`
}

func (gc *GetAccount) Execute(accountId string) (*Account, error) {
	account, err := gc.accountDAO.GetAccountByID(accountId)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func NewGetAccountCase(accountDAO AccountDAO) *GetAccount {
	return &GetAccount{
		accountDAO: accountDAO,
	}
}
