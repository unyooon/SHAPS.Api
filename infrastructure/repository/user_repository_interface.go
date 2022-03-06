package repository

import (
	"shaps.api/domain/exception"
	"shaps.api/entity"
)

type UserRepositoryInterface interface {
	Create(req entity.User) (entity.User, exception.Wrapper)
	Read(id string) (entity.User, exception.Wrapper)
	CreateStripeConnect(id string, connectId string) (entity.User, exception.Wrapper)
}
