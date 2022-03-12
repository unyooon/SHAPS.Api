package exception

import "go.uber.org/zap"

type CustomException struct {
	Code       Code
	Message    Message
	Err        error
	StackTrace string
}

func NewCustomException(c Code, m Message, e error) *CustomException {
	stack := zap.Stack("").String
	return &CustomException{
		Code:       c,
		Message:    m,
		Err:        e,
		StackTrace: stack,
	}
}
