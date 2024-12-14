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

var databaseSet = wire.NewSet(
	database.NewPostgresAdapter,
	wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)),
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
)

var positionRepoSet = wire.NewSet(
	repository.NewPositionRepository,
	wire.Bind(new(repository.PositionRepository), new(*repository.PositionRepositoryDatabase)),
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
		positionRepoSet,
		databaseSet,
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

func NewStartRide() *usecase.StartRide {
	wire.Build(
		usecase.NewStartRideUseCase,
		rideSet,
		databaseSet,
	)
	return &usecase.StartRide{}
}

func NewUpdatePosition() *usecase.UpdatePosition {
	wire.Build(
		usecase.NewUpdatePositionUseCase,
		rideSet,
		positionRepoSet,
		databaseSet,
	)
	return &usecase.UpdatePosition{}
}

func NewFinishRide() *usecase.FinishRide {
	wire.Build(
		usecase.NewFinishRideUseCase,
		rideSet,
		positionRepoSet,
		databaseSet,
	)
	return &usecase.FinishRide{}
}
