package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/exception"
)

type StripeConnectCreater interface {
	Execute(c *gin.Context) exception.CustomException
}
