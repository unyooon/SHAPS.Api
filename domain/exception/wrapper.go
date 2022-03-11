package exception

import "go.uber.org/zap"

type Wrapper struct {
	Code       Code
	Message    Message
	Err        error
	StackTrace string
}

func NewWrapper(c Code, m Message, e error) *Wrapper {
	stack := zap.Stack("").String
	return &Wrapper{
		Code:       c,
		Message:    m,
		Err:        e,
		StackTrace: stack,
	}
}
