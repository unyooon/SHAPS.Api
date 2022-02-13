package exception

import (
	"log"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Wrapper struct {
	Code    Code
	Message Message
	Err     error
}

func NewWrapper(c Code, m Message, e error) *Wrapper {
	return &Wrapper{
		Code:    c,
		Message: m,
		Err:     e,
	}
}

func (w *Wrapper) Error() {
	traceId, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Error(
		string(w.Message),
		zap.String("traceId", traceId.String()),
		zap.String("error", w.Err.Error()),
	)
}
