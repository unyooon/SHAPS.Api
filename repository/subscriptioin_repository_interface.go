package repository

import (
	"shaps.api/entity"
)

type SubscriptionRepositoryInterface interface {
	Create(req entity.Subscription) (s entity.Subscription, err error)
}
