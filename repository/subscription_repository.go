package repository

import (
	"github.com/jinzhu/gorm"
	"shaps.api/entity"
	"shaps.api/infrastructure"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db infrastructure.DbInterface) *SubscriptionRepository {
	d := db.Connect()
	return &SubscriptionRepository{
		db: d,
	}
}

func (repo *SubscriptionRepository) Create(req entity.Subscription) (s entity.Subscription, err error) {
	result := repo.db.Create(&req)

	return req, result.Error
}
