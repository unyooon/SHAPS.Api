package types

import "net/http"

type HttpRequestLogType struct {
	BaseLogType
	RequestMethod      string
	RequestHeaders     http.Header
	RequestUrl         string
	RequestQueryString string
	RequestBody        string
}
