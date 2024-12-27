package application

import (
	"github.com.br/gibranct/ride/internal/ride/application/usecase"
	di "github.com.br/gibranct/ride/internal/ride/infra/DI"
)

type RideService struct {
	*usecase.GetRide
	*usecase.RequestRide
	*usecase.StartRide
	*usecase.AcceptRide
	*usecase.FinishRide
}

type Application struct {
	*RideService
}

func NewApplication() *Application {
	return &Application{
		RideService: &RideService{
			GetRide:     di.NewGetRide(),
			RequestRide: di.NewRequestRide(),
			StartRide:   di.NewStartRide(),
			AcceptRide:  di.NewAcceptRide(),
			FinishRide:  di.NewFinishRide(),
		},
	}
}
