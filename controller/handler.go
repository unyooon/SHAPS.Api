package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"shaps.api/domain/exception"
)

func Handler(c *gin.Context, obj interface{}, err exception.Wrapper) {
	if err.Code == exception.OkCode {
		c.Set("resBody", obj)
		c.JSON(int(err.Code), obj)
		return
	} else {
		traceId, exists := c.Get("traceId")
		if !exists {
			traceIdErr := &exception.Wrapper{
				Code:    exception.InternalServerErrorCode,
				Message: exception.InternalServerErrorMessage,
				Err:     errors.New("can not get traceId"),
			}
			c.Set("err", traceIdErr)
			c.JSON(int(traceIdErr.Code), traceIdErr.Message)
		}

		c.Set("err", err)
		c.JSON(int(err.Code), &exception.Response{
			Message: string(err.Message),
			TraceId: traceId.(uuid.UUID).String(),
		})
	}
}
