//go:build wireinject
// +build wireinject

package di

import (
	"github.com.br/gibranct/ride/internal/ride/application/usecase"
	"github.com.br/gibranct/ride/internal/ride/infra/database"
	"github.com.br/gibranct/ride/internal/ride/infra/gateway"
	"github.com.br/gibranct/ride/internal/ride/infra/queue"
	"github.com.br/gibranct/ride/internal/ride/infra/repository"
	"github.com/google/wire"
)

var databaseSet = wire.NewSet(
	database.NewPostgresAdapter,
	wire.Bind(new(database.DatabaseConnection), new(*database.PostgresAdapter)),
)

var allReposSet = wire.NewSet(
	gateway.NewAccountGateway,
	repository.NewRideRepository,
	wire.Bind(new(repository.RideRepository), new(*repository.RideRepositoryDatabase)),
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
		queue.NewRabbitMQAdapter,
	)
	return &usecase.FinishRide{}
}
