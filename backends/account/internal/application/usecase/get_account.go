package usecase

import (
	"github.com.br/gibranct/account/internal/domain/errors"
	"github.com.br/gibranct/account/internal/infra/repository"
)

type GetAccount struct {
	accountRepository repository.AccountRepository
}

type GetAccountOutput struct {
	ID          string
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	IsPassenger bool
	IsDriver    bool
}

func (gc *GetAccount) Execute(accountId string) (*GetAccountOutput, error) {
	account, err := gc.accountRepository.GetAccountByID(accountId)

	if err != nil {
		return nil, errors.ErrAccountNotFound
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

func NewGetAccountUseCase(accountRepository repository.AccountRepository) *GetAccount {
	return &GetAccount{
		accountRepository: accountRepository,
	}
}
