package domain

import (
	"errors"

	"github.com/gin-gonic/gin"
	"shaps.api/entity"
	"shaps.api/repository"
)

type CreateUserInteractor struct {
	UserRepository repository.UserRepositoryInterface
}

func NewCreateUserInteractor(r repository.UserRepositoryInterface) *CreateUserInteractor {
	return &CreateUserInteractor{
		UserRepository: r,
	}
}

func (i *CreateUserInteractor) Excecute(c *gin.Context) (err error) {
	uid, exists := c.Get("userId")
	if !exists {
		return errors.New("not found userId")
	}

	u := entity.User{
		ID: uid.(string),
	}

	_, err = i.UserRepository.Create(u)

	return err
}
