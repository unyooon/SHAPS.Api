package repository

import (
	"github.com/jinzhu/gorm"
	"shaps.api/entity"
)

type SubscriptionRepository struct{}

func (repo *SubscriptionRepository) Create(d *gorm.DB, req entity.Subscription) (s entity.Subscription, err error) {
	result := d.Create(&req)

	return req, result.Error
}
