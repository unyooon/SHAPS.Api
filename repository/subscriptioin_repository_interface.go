package repository

import (
	"github.com/jinzhu/gorm"
	"shaps.api/entity"
)

type SubscriptionRepositoryInterface interface {
	Create(d *gorm.DB, req entity.Subscription) (s entity.Subscription, err error)
}
