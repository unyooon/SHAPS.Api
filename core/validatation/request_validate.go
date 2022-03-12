package validatation

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"shaps.api/domain/exception"
)

func RequestValidate(obj interface{}, c *gin.Context) exception.CustomException {
	if err := c.ShouldBindBodyWith(&obj, binding.JSON); err != nil {
		e := exception.CustomException{
			Code:    exception.BadRequestCode,
			Message: exception.BadRequestMessage,
			Err:     err,
		}
		return e
	}
	return exception.CustomException{}
}
