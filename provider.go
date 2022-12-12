package main

import (
	"github.com/google/wire"
	"shaps.api/controller"
	"shaps.api/domain"
	"shaps.api/infrastructure/db"
	"shaps.api/infrastructure/external"
	"shaps.api/infrastructure/repository"
	"shaps.api/usecase"
)

var SuperSet = wire.NewSet(
	// controller
	controller.NewSubscriptionController,
	controller.NewUserController,
	controller.NewMeController,

	// usecase
	domain.NewCreateSubscriptionInteractor,
	domain.NewCreateUserInteractor,
	domain.NewReadUserInteractor,
	domain.NewCreateStripeConnectInteractor,
	domain.NewReadMeInteractor,
	domain.NewUpdateMeInteractor,
	domain.NewReadHostsInteractor,
	domain.NewReadConstructsInteractor,
	domain.NewJoinSubscriptionInteractor,
	wire.Bind(new(usecase.SubscriptionCreater), new(*domain.CreateSubscriptionInteractor)),
	wire.Bind(new(usecase.UserCreater), new(*domain.CreateUserInteractor)),
	wire.Bind(new(usecase.UserReader), new(*domain.ReadUserInteractor)),
	wire.Bind(new(usecase.StripeConnectCreater), new(*domain.CreateStripeConnectInteractor)),
	wire.Bind(new(usecase.MeReader), new(*domain.ReadMeInteractor)),
	wire.Bind(new(usecase.MeUpdater), new(*domain.UpdateMeInteractor)),
	wire.Bind(new(usecase.HostsReader), new(*domain.ReadHostsInteractor)),
	wire.Bind(new(usecase.ConstructsReader), new(*domain.ReadConstructsInteractor)),
	wire.Bind(new(usecase.SubscriptionJoiner), new(*domain.JoinSubscriptionInteractor)),

	// repository
	repository.NewSubscriptionRepository,
	repository.NewUserRepository,
	wire.Bind(new(repository.SubscriptionRepositoryInterface), new(*repository.SubscriptionRepository)),
	wire.Bind(new(repository.UserRepositoryInterface), new(*repository.UserRepository)),

	//client
	external.NewStripeClient,

	// db
	db.NewDb,
	wire.Bind(new(db.DbInterface), new((*db.Db))),
)
