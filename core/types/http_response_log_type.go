package types

import "net/http"

type HttpResponseLogType struct {
	BaseLogType
	Time               string
	ResponseHeaders    http.Header
	ResponseStatusCode string
	ResponseBody       string
}
