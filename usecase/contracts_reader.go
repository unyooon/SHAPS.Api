package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
)

type ContractsReader interface {
	Execute(c *gin.Context) (dto.ReadContractsResponse, *exception.CustomException)
}
