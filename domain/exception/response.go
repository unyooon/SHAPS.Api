package exception

type Response struct {
	Message string `column:"message"`
	TraceId string `column:"traceId"`
}
