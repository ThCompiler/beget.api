package response

import "github.com/pkg/errors"

var (
	ErrAPIMethodReturnError   = errors.New("method was executed without success result, it's return error")
	ErrAPIMethodReturnSuccess = errors.New("method was executed with success result, it's not return error")
	ErrAPIReturnError         = errors.New("response was executed without success result, it's return error")
	ErrAPIReturnSuccess       = errors.New("response was executed with success result, it's not return error")
)
