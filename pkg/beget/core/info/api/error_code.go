package api

type ErrorCode string

const (
	AuthError        = ErrorCode("AUTH_ERROR")
	IncorrectRequest = ErrorCode("INCORRECT_REQUEST")
	NoSuchMethod     = ErrorCode("NO_SUCH_METHOD")
)
