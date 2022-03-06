package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/exception"
)

type StripeConnectCreater interface {
	Excute(c *gin.Context) exception.Wrapper
}
