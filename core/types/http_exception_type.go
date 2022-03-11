package types

type HttpExceptionType struct {
	HttpStatusCode string
	ErrorMessage   string
	Source         string
	StackTrace     string
}
