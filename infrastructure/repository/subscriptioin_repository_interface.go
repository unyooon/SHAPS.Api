package repository

import (
	"shaps.api/domain/exception"
	"shaps.api/entity"
)

type SubscriptionRepositoryInterface interface {
	Create(user entity.User, subscription entity.Subscription) (entity.Subscription, *exception.CustomException)
	ReadHosts(user entity.User) ([]entity.Subscription, *exception.CustomException)
	ReadContracts(user entity.User) ([]entity.Contract, *exception.CustomException)
  ReadContract(user entity.User, id uint) (entity.Contract, *exception.CustomException)
	JoinSubscription(user entity.User, subscription entity.Subscription) (*exception.CustomException)
	ReadSubscription(id uint) (entity.Subscription, *exception.CustomException)
	CancelContract(contract entity.Contract) *exception.CustomException
}
