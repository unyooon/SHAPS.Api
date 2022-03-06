package middleware

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"shaps.api/core/constants"
	"shaps.api/core/logger"
	"shaps.api/core/types"
)

func HttpLogger(c *gin.Context) {
	start := time.Now()

	traceId, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err.Error())
	}

	hostname, _ := os.Hostname()
	uid, _ := c.Get("userId")
	reqTs := start.UTC().Format("2006-01-02T15:04:05+09:00")
	reqBody := make([]byte, c.Request.ContentLength)
	c.Request.Body.Read(reqBody)

	requestLog := types.HttpRequestLogType{
		BaseLogType: types.BaseLogType{
			ApplicationName: "SHAPS.Api",
			HostName:        hostname,
			Placement:       "shaps.api/middleware/http_logger",
			UserId:          uid.(string),
			TraceId:         traceId.String(),
			TimeStamp:       reqTs,
		},
		RequestHeaders:     c.Request.Header,
		RequestUrl:         c.Request.URL.Path,
		RequestQueryString: c.Request.URL.RawQuery,
		RequestBody:        string(reqBody),
	}
	logger.Logger(requestLog, constants.HttpRequestLog)

	c.Next()

	resTs := time.Now().UTC().Format("2006-01-02T15:04:05+09:00")
	end := time.Since(start).Milliseconds()
	resHeader := c.Writer.Header()
	statusCode := c.Writer.Status()
	resBody, _ := c.Get("resBody")

	responseLog := types.HttpResponseLogType{
		BaseLogType: types.BaseLogType{
			ApplicationName: "SHAPS.Api",
			HostName:        hostname,
			Placement:       "shaps.api/middleware/http_logger",
			UserId:          uid.(string),
			TraceId:         traceId.String(),
			TimeStamp:       resTs,
		},
		Time:               strconv.FormatInt(end, 10),
		ResponseHeaders:    resHeader,
		ResponseStatusCode: strconv.FormatInt(int64(statusCode), 10),
		ResponseBody:       resBody.(string),
	}
	logger.Logger(responseLog, constants.HttpResponseLog)
}
