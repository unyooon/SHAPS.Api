package domain

import (
	"errors"

	"github.com/gin-gonic/gin"
	"shaps.api/core/constants"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/repository"
)

type Query struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type ReadHostsInteractor struct {
	UserRepository repository.UserRepositoryInterface
	SubscriptionRepository repository.SubscriptionRepositoryInterface
}

func NewReadHostsInteractor(
	ur repository.UserRepositoryInterface,
	sr repository.SubscriptionRepositoryInterface,
) *ReadHostsInteractor {
	return &ReadHostsInteractor{
		UserRepository: ur,
		SubscriptionRepository: sr,
	}
}

func (i *ReadHostsInteractor) Execute(c *gin.Context) (dto.ReadHostsResponse, *exception.CustomException) {
	uid, exists := c.Get("userId")
	if !exists {
		e := &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		return dto.ReadHostsResponse{}, e
	}

	var query Query
	if err := c.ShouldBindQuery(&query); err != nil {
		e := &exception.CustomException{
			Code: constants.BadRequestCode,
			Message: constants.BadRequestMessage,
			Err: err,
		}
		return dto.ReadHostsResponse{}, e
	}


	if query.Size <= 0 || query.Size > 50 {
		query.Size = 50
	}
	if query.Page <= 0 {
		query.Page = 1
	}

	u, userErr := i.UserRepository.Read(uid.(string))
	if userErr != nil {
		return dto.ReadHostsResponse{}, userErr
	}

	hosts, hostsErr := i.SubscriptionRepository.ReadHosts(u)
	if hostsErr != nil {
		return dto.ReadHostsResponse{}, userErr
	}

	if len(hosts) / query.Size + 1 < query.Page {
		e := &exception.CustomException{
			Code: constants.NotFoundCode,
			Message: constants.NotFoundMessage,
			Err: errors.New("the page does not exist"),
		}
		return dto.ReadHostsResponse{}, e
	}

	var data []dto.ReadHostResponse
	for i:=0; i<len(hosts); i++ {
		data = append(data, dto.ReadHostResponse{
			ID: hosts[i].ID,
			Name: hosts[i].Name,
			Price: hosts[i].Price,
			Term: hosts[i].Term,
			Description: hosts[i].Description,
			CreatedAt: hosts[i].CreatedAt.Format("2006-01-02T15:04:05+09:00"),
		})
	}

	return dto.ReadHostsResponse{
		Page: dto.BasePageResponse{
			Number: query.Page,
			Size: query.Size,
			TotalElements: len(hosts),
			TotalPages: len(hosts) / query.Size + 1,
		},
		Data: data,
	}, nil
}
