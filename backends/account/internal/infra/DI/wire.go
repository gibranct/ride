//go:build wireinject
// +build wireinject

package di

import (
	"github.com.br/gibranct/account/internal/application/usecase"
	"github.com.br/gibranct/account/internal/infra/database"
	"github.com.br/gibranct/account/internal/infra/gateway"
	"github.com.br/gibranct/account/internal/infra/repository"
	"github.com/google/wire"
)

var databaseSet = wire.NewSet(
	database.NewPostgresAdapter,
	wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)),
)

var accountSet = wire.NewSet(
	repository.NewAccountRepository,
	wire.Bind(new(repository.AccountRepository), new(*repository.AccountRepositoryDatabase)),
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

func NewAccountPostgresRepository() *repository.AccountRepositoryDatabase {
	wire.Build(
		repository.NewAccountRepository,
		databaseSet,
	)
	return &repository.AccountRepositoryDatabase{}
}
