package usecase

import (
	"fmt"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com.br/gibranct/ride/cmd/infra/gateway"
	"github.com.br/gibranct/ride/cmd/infra/repository"
)

type SignUp struct {
	accountDAO    repository.AccountRepository
	mailerGateway gateway.MailerGateway
}

type SignUpOutput struct {
	AccountId string `json:"accountId"`
}

type SignUpInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
	Password    string `json:"password"`
}

func (signUp *SignUp) Execute(input SignUpInput) (*SignUpOutput, error) {
	newAccount, err := domain.CreateAccount(
		input.Name, input.Email, input.CPF, input.CarPlate,
		input.Password, input.IsPassenger, input.IsDriver,
	)

	if err != nil {
		return nil, err
	}

	account, err := signUp.accountDAO.GetAccountByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if account.ID != "" {
		return nil, fmt.Errorf("duplicated account")
	}

	err = signUp.accountDAO.SaveAccount(*newAccount)

	if err != nil {
		return nil, err
	}

	signUp.mailerGateway.Send(newAccount.GetEmail(), "Welcome!", "...")

	return &SignUpOutput{
		AccountId: newAccount.ID,
	}, nil
}

func NewSignUpUseCase(accountDAO repository.AccountRepository, mailer gateway.MailerGateway) *SignUp {
	return &SignUp{
		accountDAO:    accountDAO,
		mailerGateway: mailer,
	}
}
