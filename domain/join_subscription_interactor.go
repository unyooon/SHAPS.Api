package domain

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"shaps.api/core/constants"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/repository"
)

type JoinSubscriptionInteractor struct {
	SubscriptionRepository repository.SubscriptionRepositoryInterface
	UserRepository         repository.UserRepositoryInterface
}

func NewJoinSubscriptionInteractor(sr repository.SubscriptionRepositoryInterface, ur repository.UserRepositoryInterface) *JoinSubscriptionInteractor {
	return &JoinSubscriptionInteractor{
		SubscriptionRepository: sr,
		UserRepository: ur,
	}
}

func (i *JoinSubscriptionInteractor) Execute(c *gin.Context) *exception.CustomException {
	uid, exists := c.Get("userId")
	if !exists {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return e
	}

	u, userErr := i.UserRepository.Read(uid.(string))
	if userErr != nil {
		return userErr
	}

	subscriptionId, convErr := strconv.ParseUint(c.Param("subscriptionId"), 10, 64)
	if convErr != nil {
		e := &exception.CustomException{
			Code: constants.BadRequestCode,
			Message: constants.BadRequestMessage,
			Err: convErr,
		}
		return e
	}

	s, subscriptionErr := i.SubscriptionRepository.ReadSubscription(uint(subscriptionId))
	if subscriptionErr != nil {
		return subscriptionErr
	}

	e := i.SubscriptionRepository.JoinSubscription(u, s)

	return e
}