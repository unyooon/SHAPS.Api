package logger

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
	"shaps.api/core/constants"
	"shaps.api/core/types"
)

func Logger(obj interface{}, logType int) {
	logger, zapErr := zap.NewProduction()
	if zapErr != nil {
		log.Printf(zapErr.Error())
	}

	if logType == constants.HttpRequestLog {
		log, ok := obj.(types.HttpRequestLogType)
		if !ok {
			logger.Error(
				"AssertionError",
				zap.String("Message", "Can not assertion types.HttpRequestLogType"),
			)
		}

		headers, headerJsonErr := json.Marshal(log.RequestHeaders)
		if headerJsonErr != nil {
			logger.Error(
				"JsonMarshalError",
				zap.String("Message", "Can not Marshal Header Object"),
				zap.String("StackTrace", headerJsonErr.Error()),
			)
		}

		logger.Info(
			"RequestLog",
			zap.String("ApplicationName", log.ApplicationName),
			zap.String("HostName", log.HostName),
			zap.String("Placement", log.Placement),
			zap.String("UserId", log.UserId),
			zap.String("TraceId", log.TraceId),
			zap.String("TimeStamp", log.TimeStamp),
			zap.String("RequestMethod", log.RequestMethod),
			zap.String("RequestHeaders", string(headers)),
			zap.String("RequestUrl", string(log.RequestUrl)),
			zap.String("RequestQueryString", log.RequestQueryString),
			zap.String("RequestBody", log.RequestBody),
		)
	} else if logType == constants.HttpResponseLog {
		log, ok := obj.(types.HttpResponseLogType)
		if !ok {
			logger.Error(
				"AssertionError",
				zap.String("Message", "Can not assertion types.HttpResponseLogType"),
			)
		}

		headers, headerJsonErr := json.Marshal(log.ResponseHeaders)
		if headerJsonErr != nil {
			logger.Error(
				"JsonMarshalError",
				zap.String("Message", "Can not Marshal Header Object"),
				zap.String("StackTrace", headerJsonErr.Error()),
			)
		}

		logger.Info(
			"ResponseLog",
			zap.String("ApplicationName", log.ApplicationName),
			zap.String("HostName", log.HostName),
			zap.String("Placement", log.Placement),
			zap.String("UserId", log.UserId),
			zap.String("TraceId", log.TraceId),
			zap.String("TimeStamp", log.TimeStamp),
			zap.String("Time", log.Time),
			zap.String("ResponseHeaders", string(headers)),
			zap.String("ResponseStatusCode", log.ResponseStatusCode),
			zap.String("ResponseBody", log.ResponseBody),
		)
	} else if logType == constants.HttpExceptionResponseLog {
		log, ok := obj.(types.HttpExceptionLogType)
		if !ok {
			logger.Error(
				"AssertionError",
				zap.String("Message", "Can not assertion types.HttpExceptionLogType"),
			)
		}

		logger.Error(
			"ErrorLog",
			zap.String("ApplicationName", log.ApplicationName),
			zap.String("HostName", log.HostName),
			zap.String("Placement", log.Placement),
			zap.String("UserId", log.UserId),
			zap.String("TraceId", log.TraceId),
			zap.String("TimeStamp", log.TimeStamp),
			zap.String("HttpStatusCode", log.HttpStatusCode),
			zap.String("ErrorMessage", log.ErrorMessage),
			zap.String("Source", log.Source),
			zap.String("StackTrace", log.StackTrace),
		)
	}
}
