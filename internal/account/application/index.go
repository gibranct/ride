package application

import (
	"github.com.br/gibranct/ride/internal/account/application/usecase"
	di "github.com.br/gibranct/ride/internal/account/infra/DI"
)

type AccountService struct {
	*usecase.SignUp
	*usecase.GetAccount
}

type Application struct {
	*AccountService
}

func NewApplication() *Application {
	return &Application{
		AccountService: &AccountService{
			GetAccount: di.NewGetAccount(),
			SignUp:     di.NewSignUp(),
		},
	}
}
