package usecase

import (
	"log"

	domain "github.com.br/gibranct/account/internal/domain/entity"
	"github.com.br/gibranct/account/internal/domain/errors"
	"github.com.br/gibranct/account/internal/infra/gateway"
	"github.com.br/gibranct/account/internal/infra/repository"
)

type SignUp struct {
	accountDAO    repository.AccountRepository
	mailerGateway gateway.MailerGateway
}

type SignUpOutput struct {
	AccountId string
}

type SignUpInput struct {
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	IsPassenger bool
	IsDriver    bool
	Password    string
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

	if account != nil && account.ID != "" {
		return nil, errors.ErrEmailAlreadyTaken
	}

	err = signUp.accountDAO.SaveAccount(*newAccount)

	if err != nil {
		log.Default().Println(err)
		return nil, errors.ErrSavingAccount
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
