// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com.br/gibranct/account/internal/application/usecase"
	"github.com.br/gibranct/account/internal/infra/database"
	"github.com.br/gibranct/account/internal/infra/gateway"
	"github.com.br/gibranct/account/internal/infra/repository"
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

func NewAccountPostgresRepository() *repository.AccountRepositoryDatabase {
	postgresAdapter := database.NewPostgresAdapter()
	accountRepositoryDatabase := repository.NewAccountRepository(postgresAdapter)
	return accountRepositoryDatabase
}

// wire.go:

var databaseSet = wire.NewSet(database.NewPostgresAdapter, wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)))

var accountSet = wire.NewSet(repository.NewAccountRepository, wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryDatabase)), database.NewPostgresAdapter, wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)))
