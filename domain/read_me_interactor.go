package domain

import (
	"errors"

	"github.com/gin-gonic/gin"
	"shaps.api/core/constants"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/repository"
)

type ReadMeInteractor struct {
	UserRepository repository.UserRepositoryInterface
}

func NewReadMeInteractor(
	r repository.UserRepositoryInterface,
) *ReadMeInteractor {
	return &ReadMeInteractor{
		UserRepository: r,
	}
}

func (i *ReadMeInteractor) Execute(c *gin.Context) (dto.ReadMeResponse, *exception.CustomException) {
	uid, exists := c.Get("userId")
	if !exists {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return dto.ReadMeResponse{}, e
	}

	u, err := i.UserRepository.Read(uid.(string))
	return dto.ReadMeResponse{
		ID:   u.ID,
		Icon: u.Icon,
	}, err
}
