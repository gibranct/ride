package application

import (
	"github.com.br/gibranct/ride/cmd/application/usecase"
	di "github.com.br/gibranct/ride/cmd/infra/DI"
)

type AccountService struct {
	*usecase.SignUp
	*usecase.GetAccount
}

type RideService struct {
	*usecase.GetRide
	*usecase.RequestRide
}

type Application struct {
	*AccountService
	*RideService
}

func NewApplication() *Application {
	return &Application{
		AccountService: &AccountService{
			GetAccount: di.NewGetAccount(),
			SignUp:     di.NewSignUp(),
		},
		RideService: &RideService{
			GetRide:     di.NewGetRide(),
			RequestRide: di.NewRequestRide(),
		},
	}
}
