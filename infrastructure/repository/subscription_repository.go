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

func (repo *SubscriptionRepository) ReadHosts(user entity.User) ([]entity.Subscription, *exception.CustomException) {
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

func (repo *SubscriptionRepository) ReadSubscription(id uint) (entity.Subscription, *exception.CustomException) {

	var s entity.Subscription
	result := repo.db.First(&s, "id = ?", id)
	if result.Error != nil {
		e := &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     result.Error,
		}
		return s, e
	}

	return s, nil
}

func (repo *SubscriptionRepository) JoinSubscription(user entity.User, subscription entity.Subscription) (*exception.CustomException) {
	if err:= repo.db.Model(&user).Association("Construct").Append(&entity.Construct{SubscriptionID: subscription.ID}).Error; err != nil {
		return &exception.CustomException{
			Code: constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err: err,
		}
	}

	return nil
}
