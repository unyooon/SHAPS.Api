package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"shaps.api/core/constants"
	"shaps.api/domain/exception"
)

func Handler(c *gin.Context, obj interface{}, err *exception.CustomException) {
	if err != nil {
		traceId, exists := c.Get("traceId")
		if !exists {
			traceIdErr := &exception.CustomException{
				Code:    constants.InternalServerErrorCode,
				Message: constants.InternalServerErrorMessage,
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
	} else {
		c.Set("resBody", obj)
		c.JSON(200, obj)
		return
	}
}
