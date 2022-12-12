package repository

import (
	"shaps.api/domain/exception"
	"shaps.api/entity"
)

type SubscriptionRepositoryInterface interface {
	Create(user entity.User, subscription entity.Subscription) (entity.Subscription, *exception.CustomException)
	ReadHosts(user entity.User) ([]entity.Subscription, *exception.CustomException)
	ReadConstructs(user entity.User) ([]entity.Construct, *exception.CustomException)
	JoinSubscription(user entity.User, subscription entity.Subscription) (*exception.CustomException)
	ReadSubscription(id uint) (entity.Subscription, *exception.CustomException)
}
