package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
)

type ConstructsReader interface {
	Execute(c *gin.Context) (dto.ReadConstructsResponse, *exception.CustomException)
}
