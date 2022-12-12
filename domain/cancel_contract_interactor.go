package domain

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"shaps.api/core/constants"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/repository"
)

type CancelContractInteractor struct {
	SubscriptionRepository repository.SubscriptionRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func NewCancelContractInteractor(sr repository.SubscriptionRepositoryInterface, ur repository.UserRepositoryInterface) *CancelContractInteractor {
	return &CancelContractInteractor{
		SubscriptionRepository: sr,
		UserRepository: ur,
	}
}

func (i *CancelContractInteractor) Execute(c *gin.Context) *exception.CustomException {
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

	contractId, convErr := strconv.ParseUint(c.Param("contractId"), 10, 64)
	if convErr != nil {
		e := &exception.CustomException{
			Code: constants.BadRequestCode,
			Message: constants.BadRequestMessage,
			Err: convErr,
		}
		return e
	}

	contract, contractErr := i.SubscriptionRepository.ReadContract(u, uint(contractId))
	if contractErr != nil {
		return contractErr
	}

	if err := i.SubscriptionRepository.CancelContract(contract); err != nil {
		return err
	}

	return nil
}