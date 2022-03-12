package domain

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"shaps.api/core/validatation"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/entity"
	"shaps.api/infrastructure/repository"
)

type CreateSubscriptionInteractor struct {
	SubscriptionRepository repository.SubscriptionRepositoryInterface
}

func NewCreateSubscriptionInteractor(r repository.SubscriptionRepositoryInterface) *CreateSubscriptionInteractor {
	return &CreateSubscriptionInteractor{
		SubscriptionRepository: r,
	}
}

func (i *CreateSubscriptionInteractor) Execute(c *gin.Context) exception.CustomException {
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	req := new(dto.CreateSubscriptionRequest)
	ve := validatation.RequestValidate(req, c)
	if ve.Code == exception.BadRequestCode {
		return ve
	}
	json.Unmarshal(body, &req)

	s := entity.Subscription{
		Name:        req.Name,
		Price:       req.Price,
		Term:        req.Term,
		Description: req.Description,
	}

	_, err := i.SubscriptionRepository.Create(s)

	return err
}
