package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"shaps.api/core/constants"
	"shaps.api/core/validatation"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/repository"
)

type UpdateMeInteractor struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUpdateMeInteractor(
	ur repository.UserRepositoryInterface,
) *UpdateMeInteractor {
	return &UpdateMeInteractor{
		UserRepository: ur,
	}
}

func (i *UpdateMeInteractor) Execute(c *gin.Context) *exception.CustomException {
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	req := new(dto.UpdateMeRequest)
	if ve := validatation.RequestValidate(req, c); ve != nil {
		return ve
	}

	json.Unmarshal(body, &req)

	uid, exists := c.Get("userId")
	if !exists {
		return &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUser,
			Err:     errors.New("userId is not exists"),
		}
	}

	user, readErr := i.UserRepository.Read(uid.(string))
	if readErr != nil {
		return readErr
	}

	user.Name = &req.Name
	user.Icon = &req.Icon

	if _, updateErr := i.UserRepository.Update(user); updateErr != nil {
		return updateErr
	}

	return nil
}
