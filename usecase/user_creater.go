package usecase

import "github.com/gin-gonic/gin"

type UserCreater interface {
	Excecute(c *gin.Context) (err error)
}
