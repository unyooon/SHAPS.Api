package repository

import (
	"shaps.api/domain/exception"
	"shaps.api/entity"
)

type SubscriptionRepositoryInterface interface {
	Create(req entity.Subscription) (entity.Subscription, exception.Wrapper)
}
