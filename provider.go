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
	controller.NewMeController,

	// usecase
	domain.NewCreateSubscriptionInteractor,
	domain.NewCreateUserInteractor,
	wire.Bind(new(usecase.SubscriptionCreater), new(*domain.CreateSubscriptionInteractor)),
	wire.Bind(new(usecase.UserCreater), new(*domain.CreateUserInteractor)),

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
