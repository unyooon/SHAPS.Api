package validatation

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"shaps.api/domain/exception"
)

func RequestValidate(obj interface{}, c *gin.Context) exception.Wrapper {
	if err := c.ShouldBindBodyWith(&obj, binding.JSON); err != nil {
		e := exception.Wrapper{
			Code:    exception.BadRequestCode,
			Message: exception.BadRequestMessage,
			Err:     err,
		}
		return e
	}
	return exception.Wrapper{}
}
