package domain

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/exception"
	"shaps.api/entity"
	"shaps.api/infrastructure/repository"
)

type CreateUserInteractor struct {
	UserRepository repository.UserRepositoryInterface
}

func NewCreateUserInteractor(r repository.UserRepositoryInterface) *CreateUserInteractor {
	return &CreateUserInteractor{
		UserRepository: r,
	}
}

func (i *CreateUserInteractor) Excecute(c *gin.Context) exception.Wrapper {
	uid, exists := c.Get("userId")
	if !exists {
		return exception.Wrapper{
			Code:    exception.NotFoundCode,
			Message: exception.NotFoundUserId,
		}
	}

	u := entity.User{
		ID: uid.(string),
	}

	_, err := i.UserRepository.Create(u)

	return err
}
