package repository

import (
	"github.com/jinzhu/gorm"
	"shaps.api/domain/exception"
	"shaps.api/entity"
	"shaps.api/infrastructure/db"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db db.DbInterface) *SubscriptionRepository {
	d := db.Connect()
	return &SubscriptionRepository{
		db: d,
	}
}

func (repo *SubscriptionRepository) Create(req entity.Subscription) (entity.Subscription, *exception.CustomException) {
	err := repo.db.Create(&req).Error
	if err != nil {
		e := &exception.CustomException{
			Code:    exception.InternalServerErrorCode,
			Message: exception.DatabaseError,
			Err:     err,
		}
		return req, e
	}

	return req, nil
}
