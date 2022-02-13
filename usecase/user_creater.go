package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/exception"
)

type UserCreater interface {
	Excecute(c *gin.Context) exception.Wrapper
}
