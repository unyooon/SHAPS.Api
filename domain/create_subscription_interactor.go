package domain

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/entity"
	"shaps.api/repository"
)

type CreateSubscriptionInteractor struct {
	SubscriptionRepository repository.SubscriptionRepositoryInterface
}

func NewCreateSubscriptionInteractor(r repository.SubscriptionRepositoryInterface) *CreateSubscriptionInteractor {
	return &CreateSubscriptionInteractor{
		SubscriptionRepository: r,
	}
}

func (i *CreateSubscriptionInteractor) Excecute(c *gin.Context) exception.Wrapper {
	body := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(body)

	req := new(dto.CreateSubscriptionRequest)
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
