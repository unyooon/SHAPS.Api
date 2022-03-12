package domain

import (
	"errors"

	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/repository"
)

type ReadUserInteractor struct {
	UserRepository repository.UserRepositoryInterface
}

func NewReadUserInteractor(
	r repository.UserRepositoryInterface,
) *ReadUserInteractor {
	return &ReadUserInteractor{
		UserRepository: r,
	}
}

func (i *ReadUserInteractor) Execute(c *gin.Context) (dto.ReadUserResponse, exception.CustomException) {
	uid, exists := c.Get("userId")
	if !exists {
		e := exception.CustomException{
			Code:    exception.NotFoundCode,
			Message: exception.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return dto.ReadUserResponse{}, e
	}

	u, err := i.UserRepository.Read(uid.(string))
	return dto.ReadUserResponse{
		ID:   u.ID,
		Icon: u.Icon,
	}, err
}
