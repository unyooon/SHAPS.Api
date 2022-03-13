package exception

import (
	"go.uber.org/zap"
	"shaps.api/core/constants"
)

type CustomException struct {
	Code       constants.Code
	Message    constants.Message
	Err        error
	StackTrace string
}

func NewCustomException(c constants.Code, m constants.Message, e error) *CustomException {
	stack := zap.Stack("").String
	return &CustomException{
		Code:       c,
		Message:    m,
		Err:        e,
		StackTrace: stack,
	}
}
