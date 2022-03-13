package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"shaps.api/core/constants"
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
			Code:    constants.BadRequestCode,
			Message: constants.BadRequestAlreadyExistsMessage,
		}
	}

	result := repo.db.Create(&req)
	if result.Error != nil {
		e := &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
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
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     result.Error,
		}
		return u, e
	}

	return u, nil
}

func (repo *UserRepository) Update(user entity.User) (entity.User, *exception.CustomException) {
	log.Print(user)
	if err := repo.db.Save(&user).Error; err != nil {
		return user, &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.DatabaseError,
			Err:     err,
		}
	}
	return user, nil
}
