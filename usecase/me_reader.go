package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
)

type MeReader interface {
	Execute(c *gin.Context) (dto.ReadMeResponse, *exception.CustomException)
}
