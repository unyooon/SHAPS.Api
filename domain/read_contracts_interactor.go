package domain

import (
	"errors"

	"github.com/gin-gonic/gin"
	"shaps.api/core/constants"
	"shaps.api/core/helper"
	"shaps.api/core/types"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/repository"
)

type ReadContractsInteractor struct {
	UserRepository repository.UserRepositoryInterface
	SubscriptionRepository repository.SubscriptionRepositoryInterface
}

func NewReadContractsInteractor(
	ur repository.UserRepositoryInterface,
	sr repository.SubscriptionRepositoryInterface,
) *ReadContractsInteractor {
	return &ReadContractsInteractor{
		UserRepository: ur,
		SubscriptionRepository: sr,
	}
}

func (i *ReadContractsInteractor) Execute(c *gin.Context) (dto.ReadContractsResponse, *exception.CustomException) {
	uid, exists := c.Get("userId")
	if !exists {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return dto.ReadContractsResponse{}, e
	}

	var query types.Query
	query, qErr := helper.PaginationHelper(c, query, 50)
	if qErr != nil {
		return dto.ReadContractsResponse{}, qErr
	}

	u, userErr := i.UserRepository.Read(uid.(string))
	if userErr != nil {
		return dto.ReadContractsResponse{}, userErr
	}

	contracts, contractsErr := i.SubscriptionRepository.ReadContracts(u)
	if contractsErr != nil {
		return dto.ReadContractsResponse{}, userErr
	}

	if len(contracts) / query.Size + 1 < query.Page {
		e := &exception.CustomException{
			Code: constants.NotFoundCode,
			Message: constants.NotFoundMessage,
			Err: errors.New("the page does not exist"),
		}
		return dto.ReadContractsResponse{}, e
	}

	var data []dto.ReadContractResponse
	for i:=0; i<len(contracts); i++ {
		data = append(data, dto.ReadContractResponse{
			ID: contracts[i].ID,
			CreatedAt: contracts[i].CreatedAt.Format("2006-01-02T15:04:05+09:00"),
			Subscription: dto.ReadContractsSubscriptionResponse{
				ID: contracts[i].Subscription.ID,
				Name: contracts[i].Subscription.Name,
				Price: contracts[i].Subscription.Price,
				Term: contracts[i].Subscription.Term,
				Description: contracts[i].Subscription.Description,
				CreatedAt: contracts[i].Subscription.CreatedAt.Format("2006-01-02T15:04:05+09:00"),
			},
		})
	}

	return dto.ReadContractsResponse{
		Page: dto.BasePageResponse{
			Number: query.Page,
			Size: query.Size,
			TotalElements: len(contracts),
			TotalPages: len(contracts) / query.Size + 1,
		},
		Data: data,
	}, nil
}
