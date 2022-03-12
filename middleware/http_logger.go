package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"shaps.api/core/constants"
	"shaps.api/core/logger"
	"shaps.api/core/types"
	"shaps.api/domain/exception"
)

func HttpLogger(c *gin.Context) {
	start := time.Now()

	traceId, err := uuid.NewRandom()
	if err != nil {
		log.Printf(err.Error())
	}
	c.Set("traceId", traceId)

	hostname, _ := os.Hostname()
	uid, exists := c.Get("userId")
	if !exists {
		log.Printf("userId is not exists")
	}
	reqTs := start.UTC().Format("2006-01-02T15:04:05+09:00")
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

	requestLog := types.HttpRequestLogType{
		BaseLogType: types.BaseLogType{
			ApplicationName: "SHAPS.Api",
			HostName:        hostname,
			Placement:       "shaps.api/middleware/http_logger",
			UserId:          uid.(string),
			TraceId:         traceId.String(),
			TimeStamp:       reqTs,
		},
		RequestMethod:      c.Request.Method,
		RequestHeaders:     c.Request.Header,
		RequestUrl:         c.Request.URL.Path,
		RequestQueryString: c.Request.URL.RawQuery,
		RequestBody:        string(reqBody),
	}
	logger.Logger(requestLog, constants.HttpRequestLog)

	c.Next()

	statusCode := c.Writer.Status()

	e, exist := c.Get("err")
	if exist {
		httpExceptionLog := types.HttpExceptionLogType{
			BaseLogType: types.BaseLogType{
				ApplicationName: "SHAPS.Api",
				HostName:        hostname,
				Placement:       "shaps.api/middleware/http_logger",
				UserId:          uid.(string),
				TraceId:         traceId.String(),
				TimeStamp:       reqTs,
			},
			HttpExceptionType: types.HttpExceptionType{
				HttpStatusCode: strconv.FormatInt(int64(statusCode), 10),
				ErrorMessage:   string(e.(exception.Wrapper).Err.Error()),
				Source:         "SHAPS.Api",
				StackTrace:     string(e.(exception.Wrapper).StackTrace),
			},
		}
		logger.Logger(httpExceptionLog, constants.HttpExceptionResponseLog)
	}

	resTs := time.Now().UTC().Format("2006-01-02T15:04:05+09:00")
	end := time.Since(start).Milliseconds()
	resHeader := c.Writer.Header()
	resBody, _ := c.Get("resBody")
	jsonBody, _ := json.Marshal(resBody)

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
		ResponseBody:       string(jsonBody),
	}
	logger.Logger(responseLog, constants.HttpResponseLog)
}
