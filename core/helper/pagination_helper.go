package helper

import (
	"github.com/gin-gonic/gin"
	"shaps.api/core/constants"
	"shaps.api/core/types"
	"shaps.api/domain/exception"
)

func PaginationHelper(c *gin.Context, q types.Query, sizeMax int) (types.Query, *exception.CustomException) {
	query := q
	if err := c.ShouldBindQuery(&query); err != nil {
		e := &exception.CustomException{
			Code: constants.BadRequestCode,
			Message: constants.BadRequestMessage,
			Err: err,
		}
		return types.Query{}, e
	}


	if query.Size <= 0 || query.Size > sizeMax {
		query.Size = sizeMax
	}
	if query.Page <= 0 {
		query.Page = 1
	}

	return query, nil
}