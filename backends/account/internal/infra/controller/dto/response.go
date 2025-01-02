package dto

import "github.com.br/gibranct/account/internal/application/usecase"

type SignUpInputResponseDto struct {
	AccountId string `json:"account_id"`
}

func NewSignUpInputResponseDto(accountId string) *SignUpInputResponseDto {
	return &SignUpInputResponseDto{
		AccountId: accountId,
	}
}

type GetAccountResponseDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
}

func NewGetAccountResponseDto(account *usecase.GetAccountOutput) *GetAccountResponseDto {
	return &GetAccountResponseDto{
		ID:          account.ID,
		Name:        account.Name,
		Email:       account.Email,
		CPF:         account.CPF,
		CarPlate:    account.CarPlate,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}
}
