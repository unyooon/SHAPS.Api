package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"shaps.api/domain/exception"
	"shaps.api/entity"
	"shaps.api/infrastructure/external/stripeclient"
	"shaps.api/infrastructure/repository"
)

type CreateUserInteractor struct {
	UserRepository repository.UserRepositoryInterface
	StripeClient   stripeclient.Client
}

func NewCreateUserInteractor(
	r repository.UserRepositoryInterface,
	sc stripeclient.Client) *CreateUserInteractor {
	return &CreateUserInteractor{
		UserRepository: r,
		StripeClient:   sc,
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

	su, stripeErr := i.StripeClient.Customers.New(&stripe.CustomerParams{})
	if stripeErr != nil {
		return exception.Wrapper{
			Code:    exception.InternalServerErrorCode,
			Message: exception.StripeError,
		}
	}

	u := entity.User{
		ID:       uid.(string),
		StripeId: su.ID,
	}

	_, err := i.UserRepository.Create(u)

	return err
}
