package types

import "net/http"

type HttpRequestLogType struct {
	BaseLogType
	RequestHeaders     http.Header
	RequestUrl         string
	RequestQueryString string
	RequestBody        string
}
