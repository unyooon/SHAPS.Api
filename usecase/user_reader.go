package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
)

type UserReader interface {
	Excecute(c *gin.Context) (dto.ReadUserResponse, exception.Wrapper)
}
