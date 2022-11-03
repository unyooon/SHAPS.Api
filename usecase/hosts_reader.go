package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
)

type HostsReader interface {
	Execute(c *gin.Context) (dto.ReadHostsResponse, *exception.CustomException)
}
