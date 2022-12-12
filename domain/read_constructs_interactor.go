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

type ReadConstructsInteractor struct {
	UserRepository repository.UserRepositoryInterface
	SubscriptionRepository repository.SubscriptionRepositoryInterface
}

func NewReadConstructsInteractor(
	ur repository.UserRepositoryInterface,
	sr repository.SubscriptionRepositoryInterface,
) *ReadConstructsInteractor {
	return &ReadConstructsInteractor{
		UserRepository: ur,
		SubscriptionRepository: sr,
	}
}

func (i *ReadConstructsInteractor) Execute(c *gin.Context) (dto.ReadConstructsResponse, *exception.CustomException) {
	uid, exists := c.Get("userId")
	if !exists {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return dto.ReadConstructsResponse{}, e
	}

	var query types.Query
	query, qErr := helper.PaginationHelper(c, query, 50)
	if qErr != nil {
		return dto.ReadConstructsResponse{}, qErr
	}

	u, userErr := i.UserRepository.Read(uid.(string))
	if userErr != nil {
		return dto.ReadConstructsResponse{}, userErr
	}

	constructs, constructsErr := i.SubscriptionRepository.ReadConstructs(u)
	if constructsErr != nil {
		return dto.ReadConstructsResponse{}, userErr
	}

	if len(constructs) / query.Size + 1 < query.Page {
		e := &exception.CustomException{
			Code: constants.NotFoundCode,
			Message: constants.NotFoundMessage,
			Err: errors.New("the page does not exist"),
		}
		return dto.ReadConstructsResponse{}, e
	}

	var data []dto.ReadConstructResponse
	for i:=0; i<len(constructs); i++ {
		data = append(data, dto.ReadConstructResponse{
			ID: constructs[i].ID,
			CreatedAt: constructs[i].CreatedAt.Format("2006-01-02T15:04:05+09:00"),
			Subscription: dto.ReadConstructsSubscriptionResponse{
				ID: constructs[i].Subscription.ID,
				Name: constructs[i].Subscription.Name,
				Price: constructs[i].Subscription.Price,
				Term: constructs[i].Subscription.Term,
				Description: constructs[i].Subscription.Description,
				CreatedAt: constructs[i].Subscription.CreatedAt.Format("2006-01-02T15:04:05+09:00"),
			},
		})
	}

	return dto.ReadConstructsResponse{
		Page: dto.BasePageResponse{
			Number: query.Page,
			Size: query.Size,
			TotalElements: len(constructs),
			TotalPages: len(constructs) / query.Size + 1,
		},
		Data: data,
	}, nil
}
