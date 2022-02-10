package middleware

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logging(c *gin.Context) {
	// request log
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Info(
		"request",
		zap.String("path", c.Request.URL.Path),
		zap.String("query", c.Request.URL.RawQuery),
		zap.String("ip", c.ClientIP()),
		zap.String("user-agent", c.Request.UserAgent()),
	)

	c.Next()

	// response log
	logger.Info(
		"response",
		zap.String("statusCode", strconv.Itoa(c.Writer.Status())),
	)
}