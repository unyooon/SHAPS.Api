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

func (i *ReadUserInteractor) Excecute(c *gin.Context) (dto.ReadUserResponse, exception.Wrapper) {
	uid, exists := c.Get("userId")
	if !exists {
		e := exception.Wrapper{
			Code:    exception.NotFoundCode,
			Message: exception.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		e.Error()
		return dto.ReadUserResponse{}, e
	}

	u, err := i.UserRepository.Read(uid.(string))
	return dto.ReadUserResponse{
		ID:   u.ID,
		Icon: u.Icon,
	}, err
}
