package core

import (
	"github.com/pkg/errors"
)

// Errors of preparing request to Beget.API
var (
	ErrResolveBegetHost = errors.New("error resolve beget hostname : " + Host) // "invalid beget hostname"
	ErrFromAPIMethod    = errors.New("beget api method return error")          // "api method got some error"
)

// Errors of working with response of Beget.API.
var (
	ErrAPIMethodReturnError   = errors.New("method was executed without a success result, it returns error")           // "method was executed without a success result, it returns error"
	ErrAPIMethodReturnSuccess = errors.New("method was executed with a success result, it does not return error")      // "method was executed with a success result, it does not return error"
	ErrAPIReturnError         = errors.New("general API was executed without a success result, it returns error")      // "general API was executed without a success result, it returns error"
	ErrAPIReturnSuccess       = errors.New("general API was executed with a success result, it does not return error") // "general API was executed with a success result, it does not return error"
)
