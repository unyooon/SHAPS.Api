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

	// usecase
	domain.NewCreateSubscriptionInteractor,
	wire.Bind(new(usecase.SubscriptionCreater), new(*domain.CreateSubscriptionInteractor)),

	// repository
	repository.NewSubscriptionRepository,
	wire.Bind(new(repository.SubscriptionRepositoryInterface), new(*repository.SubscriptionRepository)),

	// db
	infrastructure.NewDb,
	wire.Bind(new(infrastructure.DbInterface), new((*infrastructure.Db))),
)
