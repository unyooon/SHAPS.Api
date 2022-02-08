package repository

import (
	"github.com/jinzhu/gorm"
	"shaps.api/entity"
	"shaps.api/infrastructure"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db infrastructure.DbInterface) *UserRepository {
	d := db.Connect()
	return &UserRepository{
		db: d,
	}
}

func (repo *UserRepository) Create(req entity.User) (u entity.User, err error) {
	result := repo.db.Create(&req)

	return req, result.Error
}
