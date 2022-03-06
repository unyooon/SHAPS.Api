package types

type HttpExceptionType struct {
	HttpStatusCode string
	Message        string
	Source         string
	StackTrace     string
}
