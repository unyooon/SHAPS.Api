package usecase

import (
	"github.com/gin-gonic/gin"
)

type SubscriptionCreater interface {
	Excecute(c *gin.Context) (err error)
}
