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
	"shaps.api/entity"
	"shaps.api/infrastructure/repository"
)

type CreateSubscriptionInteractor struct {
	SubscriptionRepository repository.SubscriptionRepositoryInterface
	UserRepository         repository.UserRepositoryInterface
}

func NewCreateSubscriptionInteractor(sr repository.SubscriptionRepositoryInterface, ur repository.UserRepositoryInterface) *CreateSubscriptionInteractor {
	return &CreateSubscriptionInteractor{
		SubscriptionRepository: sr,
		UserRepository: ur,
	}
}

func (i *CreateSubscriptionInteractor) Execute(c *gin.Context) *exception.CustomException {
	uid, exists := c.Get("userId")
	if !exists {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return e
	}

	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	req := new(dto.CreateSubscriptionRequest)
	if ve := validatation.RequestValidate(req, c); ve != nil {
		return ve
	}

	json.Unmarshal(body, &req)

	user, userErr := i.UserRepository.Read(uid.(string))
	if userErr != nil {
		return userErr
	}

	s := entity.Subscription{
		Name:        req.Name,
		Price:       req.Price,
		Term:        req.Term,
		Description: req.Description,
		CreatedBy:   uid.(string),
	}

	_, err := i.SubscriptionRepository.Create(user, s)

	return err
}
