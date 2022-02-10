package middleware

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Logging(c *gin.Context) {
	traceId, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err.Error())
	}

	// request log
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Info(
		"request",
		zap.String("traceId", traceId.String()),
		zap.String("path", c.Request.URL.Path),
		zap.String("query", c.Request.URL.RawQuery),
		zap.String("ip", c.ClientIP()),
		zap.String("user-agent", c.Request.UserAgent()),
	)

	c.Next()

	// response log
	logger.Info(
		"response",
		zap.String("traceId", traceId.String()),
		zap.String("statusCode", strconv.Itoa(c.Writer.Status())),
	)
}
