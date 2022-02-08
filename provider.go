package main

import (
	"github.com/google/wire"
	"shaps.api/controller"
	"shaps.api/domain"
	"shaps.api/infrastructure"
	"shaps.api/repository"
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

	// db
	infrastructure.NewDb,
	wire.Bind(new(infrastructure.DbInterface), new((*infrastructure.Db))),
)
