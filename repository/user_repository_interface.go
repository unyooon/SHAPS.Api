package repository

import "shaps.api/entity"

type UserRepositoryInterface interface {
	Create(req entity.User) (u entity.User, err error)
}
