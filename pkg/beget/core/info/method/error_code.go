package api

type ErrorCode string

const (
	InvalidData  = ErrorCode("INVALID_DATA")
	LimitError   = ErrorCode("LIMIT_ERROR")
	MethodFailed = ErrorCode("METHOD_FAILED")
)
