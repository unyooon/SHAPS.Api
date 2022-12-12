package usecase

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/exception"
)

type SubscriptionJoiner interface {
	Execute(c *gin.Context) *exception.CustomException
}
