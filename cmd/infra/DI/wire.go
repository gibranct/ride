//go:build wireinject
// +build wireinject

package di

import (
	"github.com.br/gibranct/ride/cmd/application/usecase"
	"github.com.br/gibranct/ride/cmd/infra/database"
	"github.com.br/gibranct/ride/cmd/infra/gateway"
	"github.com.br/gibranct/ride/cmd/infra/repository"
	"github.com/google/wire"
)

var allReposSet = wire.NewSet(
	repository.NewAccountRepository,
	wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryDatabase)),
	repository.NewRideRepository,
	wire.Bind(new(repository.RideRepository), new(*repository.RideRepositoryDatabase)),
	database.NewPostgresAdapter,
	wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)),
)

var accountSet = wire.NewSet(
	repository.NewAccountRepository,
	wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryDatabase)),
	database.NewPostgresAdapter,
	wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)),
)

var rideSet = wire.NewSet(
	repository.NewRideRepository,
	wire.Bind(new(repository.RideRepository), new(*repository.RideRepositoryDatabase)),
	database.NewPostgresAdapter,
	wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)),
)

func NewSignUp() *usecase.SignUp {
	wire.Build(
		usecase.NewSignUpUseCase,
		gateway.NewMailerGatewayMemory,
		accountSet,
	)
	return &usecase.SignUp{}
}

func NewGetAccount() *usecase.GetAccount {
	wire.Build(
		usecase.NewGetAccountUseCase,
		accountSet,
	)
	return &usecase.GetAccount{}
}

func NewRequestRide() *usecase.RequestRide {
	wire.Build(
		usecase.NewRequestRideUseCase,
		allReposSet,
	)
	return &usecase.RequestRide{}
}

func NewGetRide() *usecase.GetRide {
	wire.Build(
		usecase.NewGetRideUseCase,
		rideSet,
	)
	return &usecase.GetRide{}
}

func NewAcceptRide() *usecase.AcceptRide {
	wire.Build(
		usecase.NewAcceptRideUseCase,
		allReposSet,
	)
	return &usecase.AcceptRide{}
}
