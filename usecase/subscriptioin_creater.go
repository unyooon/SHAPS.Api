package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/exception"
)

type SubscriptionCreater interface {
	Execute(c *gin.Context) exception.Wrapper
}
