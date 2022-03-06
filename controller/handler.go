package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/domain/exception"
)

func Handler(c *gin.Context, obj interface{}, err exception.Wrapper) {
	if err.Code == exception.OkCode {
		c.Set("resBody", obj)
		c.JSON(int(err.Code), obj)
		return
	} else {
		c.JSON(int(err.Code), err.Message)
	}
}
