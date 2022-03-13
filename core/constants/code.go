package constants

type Code int

const (
	OkCode Code = 200

	BadRequestCode          Code = 400
	UnAuthorizedCode        Code = 401
	ForbiddenCode           Code = 403
	NotFoundCode            Code = 404
	InternalServerErrorCode Code = 500
)
