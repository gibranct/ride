package dto

import "github.com.br/gibranct/account/internal/application/usecase"

type SignUpInputRequestDto struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
	Password    string `json:"password"`
}

func (dto *SignUpInputRequestDto) ToSignUpInput() usecase.SignUpInput {
	return usecase.SignUpInput{
		Name:        dto.Name,
		Email:       dto.Email,
		CPF:         dto.CPF,
		CarPlate:    dto.CarPlate,
		IsPassenger: dto.IsPassenger,
		IsDriver:    dto.IsDriver,
		Password:    dto.Password,
	}
}
