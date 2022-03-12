package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
)

type UserReader interface {
	Execute(c *gin.Context) (dto.ReadUserResponse, *exception.CustomException)
}
