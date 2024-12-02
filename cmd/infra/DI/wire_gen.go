// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com.br/gibranct/ride/cmd/application/usecase"
	"github.com.br/gibranct/ride/cmd/infra/database"
	"github.com.br/gibranct/ride/cmd/infra/gateway"
	"github.com.br/gibranct/ride/cmd/infra/repository"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewSignUp() *usecase.SignUp {
	postgresAdapter := database.NewPostgresAdapter()
	accountRepositoryDatabase := repository.NewAccountRepository(postgresAdapter)
	mailerGateway := gateway.NewMailerGatewayMemory()
	signUp := usecase.NewSignUpUseCase(accountRepositoryDatabase, mailerGateway)
	return signUp
}

func NewGetAccount() *usecase.GetAccount {
	postgresAdapter := database.NewPostgresAdapter()
	accountRepositoryDatabase := repository.NewAccountRepository(postgresAdapter)
	getAccount := usecase.NewGetAccountUseCase(accountRepositoryDatabase)
	return getAccount
}

func NewRequestRide() *usecase.RequestRide {
	postgresAdapter := database.NewPostgresAdapter()
	accountRepositoryDatabase := repository.NewAccountRepository(postgresAdapter)
	rideRepositoryDatabase := repository.NewRideRepository(postgresAdapter)
	requestRide := usecase.NewRequestRideUseCase(accountRepositoryDatabase, rideRepositoryDatabase)
	return requestRide
}

func NewGetRide() *usecase.GetRide {
	postgresAdapter := database.NewPostgresAdapter()
	rideRepositoryDatabase := repository.NewRideRepository(postgresAdapter)
	getRide := usecase.NewGetRideUseCase(rideRepositoryDatabase)
	return getRide
}

func NewAcceptRide() *usecase.AcceptRide {
	postgresAdapter := database.NewPostgresAdapter()
	accountRepositoryDatabase := repository.NewAccountRepository(postgresAdapter)
	rideRepositoryDatabase := repository.NewRideRepository(postgresAdapter)
	acceptRide := usecase.NewAcceptRideUseCase(accountRepositoryDatabase, rideRepositoryDatabase)
	return acceptRide
}

func NewStartRide() *usecase.StartRide {
	postgresAdapter := database.NewPostgresAdapter()
	rideRepositoryDatabase := repository.NewRideRepository(postgresAdapter)
	startRide := usecase.NewStartRideUseCase(rideRepositoryDatabase)
	return startRide
}

// wire.go:

var allReposSet = wire.NewSet(repository.NewAccountRepository, wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryDatabase)), repository.NewRideRepository, wire.Bind(new(repository.RideRepository), new(*repository.RideRepositoryDatabase)), database.NewPostgresAdapter, wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)))

var accountSet = wire.NewSet(repository.NewAccountRepository, wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryDatabase)), database.NewPostgresAdapter, wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)))

var rideSet = wire.NewSet(repository.NewRideRepository, wire.Bind(new(repository.RideRepository), new(*repository.RideRepositoryDatabase)), database.NewPostgresAdapter, wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)))
