package repository

import (
	"github.com/jinzhu/gorm"
	"shaps.api/core/constants"
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

func (repo *SubscriptionRepository) Create(user entity.User, subscription entity.Subscription) (entity.Subscription, *exception.CustomException) {
	if err := repo.db.Model(&user).Association("Host").Append(subscription).Error; err != nil {
		return subscription, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return subscription, nil
}

func (repo *SubscriptionRepository) 	ReadHosts(user entity.User) ([]entity.Subscription, *exception.CustomException) {
	var hosts []entity.Subscription
	if err := repo.db.Model(&user).Association("Host").Find(&hosts).Error; err != nil {
		return hosts, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return hosts, nil
}
