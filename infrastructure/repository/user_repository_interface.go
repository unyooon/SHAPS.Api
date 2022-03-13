package repository

import (
	"shaps.api/domain/exception"
	"shaps.api/entity"
)

type UserRepositoryInterface interface {
	Create(req entity.User) (entity.User, *exception.CustomException)
	Read(id string) (entity.User, *exception.CustomException)
	Update(user entity.User) (entity.User, *exception.CustomException)
}
