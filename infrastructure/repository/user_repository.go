package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"shaps.api/domain/exception"
	"shaps.api/entity"
	"shaps.api/infrastructure/db"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db db.DbInterface) *UserRepository {
	d := db.Connect()
	return &UserRepository{
		db: d,
	}
}

func (repo *UserRepository) Create(req entity.User) (entity.User, *exception.CustomException) {
	var u entity.User
	err := repo.db.First(&u).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return u, &exception.CustomException{
			Code:    exception.BadRequestCode,
			Message: exception.BadRequestAlreadyExistsMessage,
		}
	}

	result := repo.db.Create(&req)
	if result.Error != nil {
		e := &exception.CustomException{
			Code:    exception.InternalServerErrorCode,
			Message: exception.DatabaseError,
			Err:     result.Error,
		}
		return req, e
	}

	return req, nil
}

func (repo *UserRepository) Read(id string) (entity.User, *exception.CustomException) {
	var u entity.User
	result := repo.db.First(&u, "id = ?", id)
	if result.Error != nil {
		e := &exception.CustomException{
			Code:    exception.InternalServerErrorCode,
			Message: exception.DatabaseError,
			Err:     result.Error,
		}
		return u, e
	}

	return u, nil
}

func (repo *UserRepository) CreateStripeConnect(id string, connectId string) (entity.User, *exception.CustomException) {
	var u entity.User
	findResult := repo.db.First(&u, "id = ?", id)
	if findResult.Error != nil {
		e := &exception.CustomException{
			Code:    exception.InternalServerErrorCode,
			Message: exception.DatabaseError,
			Err:     findResult.Error,
		}
		return u, e
	}

	u.ConnectId = connectId
	updateResult := repo.db.Save(&u)
	if updateResult.Error != nil {
		e := &exception.CustomException{
			Code:    exception.InternalServerErrorCode,
			Message: exception.DatabaseError,
			Err:     findResult.Error,
		}
		return u, e
	}

	return u, nil
}
