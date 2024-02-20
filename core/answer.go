package core

import (
	"encoding/json"
	"github.com/ThCompiler/go.beget.api/pkg/slices"
	"strings"
)

// MethodErrors is the type to represent an API method call method's errors.
type MethodErrors []Error[MethodErrorCode]

// Error represents MethodErrors as a string
func (m MethodErrors) Error() string {
	errorStrings := slices.Map(m, func(e Error[MethodErrorCode]) string { return e.Error() })

	return "method return errors: " + strings.Join(errorStrings, "; ")
}

// Answer defines the response of the specific method part of the API.
type Answer[Result any] struct {
	status Status
	errors MethodErrors
	result *Result
}

// Get returns the result of the API method call if the API method did not fail,
// or returns the API method errors converted to a Golang error.
func (a *Answer[Result]) Get() (*Result, error) {
	if a.errors == nil {
		return a.result, nil
	}
	return a.result, a.errors
}

// Status returns the response status of a specific method.
func (a *Answer[Result]) Status() Status {
	return a.status
}

// IsError reports whether a specific method returns some error.
func (a *Answer[Result]) IsError() bool {
	return a.errors != nil
}

// GetResult returns the result of the API method call if the API method did not fail
// or [ErrAPIMethodReturnError] if the API method failed.
func (a *Answer[Result]) GetResult() (*Result, error) {
	if a.IsError() {
		return nil, ErrAPIMethodReturnError
	}

	return a.result, nil
}

// GetError returns the result of the API method call if the API method failed
// or [ErrAPIMethodReturnSuccess] if the API method did not fail.
func (a *Answer[Result]) GetError() (MethodErrors, error) {
	if !a.IsError() {
		return nil, ErrAPIMethodReturnSuccess
	}

	return a.errors, nil
}

// MustGetResult returns the result of the API method call if the API method did not fail
// or panics with [ErrAPIMethodReturnError] if the API method failed.
func (a *Answer[Result]) MustGetResult() *Result {
	if a.IsError() {
		panic(ErrAPIMethodReturnError)
	}

	return a.result
}

// MustGetError returns the result of the API method call if the API method failed.
// or panics with [ErrAPIMethodReturnSuccess] if the API method did not fail.
func (a *Answer[Result]) MustGetError() MethodErrors {
	if !a.IsError() {
		panic(ErrAPIMethodReturnSuccess)
	}

	return a.errors
}

// jsonAnswer is a temporary structure to parse the specific method part of Beget.API response from json.
type jsonAnswer[Result any] struct {
	Status Status       `json:"status"`
	Errors MethodErrors `json:"errors,omitempty"`
	Result *Result      `json:"result,omitempty"`
}

// UnmarshalJSON is functions for [encoding/json] to unmarshal response from json format.
func (a *Answer[Result]) UnmarshalJSON(data []byte) error {
	var answer jsonAnswer[Result]
	if err := json.Unmarshal(data, &answer); err != nil {
		return err
	}

	a.status = answer.Status
	a.result = answer.Result
	a.errors = answer.Errors

	return nil
}
