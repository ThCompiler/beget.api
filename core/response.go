package core

import (
	"encoding/json"
)

// Error is the type to represent an API method call general's errors.
type Error[Err MethodErrorCode | APIErrorCode] struct {
	ErrorText string `json:"error_text"`
	ErrorCode Err    `json:"error_code"`
}

// Error represents Error as a string
func (e *Error[Err]) Error() string {
	return string(e.ErrorCode) + ": " + e.ErrorText
}

// BegetResponse defines the response of the general part of the API.
type BegetResponse[Result any] struct {
	status Status
	answer *Answer[Result]
	error  *Error[APIErrorCode]
}

// Get returns the response of the API method call if the general API did not fail,
// or returns the general API error converted to a Golang error.
func (a *BegetResponse[Result]) Get() (*Answer[Result], error) {
	if a.error == nil {
		return a.answer, nil
	}
	return a.answer, a.error
}

// Status returns the response status of the general API.
func (a *BegetResponse[Result]) Status() Status {
	return a.status
}

// IsError reports whether the general API returns some error.
func (a *BegetResponse[Result]) IsError() bool {
	return a.error != nil
}

// GetAnswer returns the response of the API method call if the general API did not fail
// or [ErrAPIReturnError] if the general API failed.
func (a *BegetResponse[Result]) GetAnswer() (*Answer[Result], error) {
	if a.IsError() {
		return nil, ErrAPIReturnError
	}

	return a.answer, nil
}

// GetError returns the response of the API method call if the general API failed
// or [ErrAPIReturnSuccess] if the general API did not fail.
func (a *BegetResponse[Result]) GetError() (*Error[APIErrorCode], error) {
	if !a.IsError() {
		return nil, ErrAPIReturnSuccess
	}

	return a.error, nil
}

// MustGetAnswer returns the response of the API method call if the general API did not fail
// or panics with [ErrAPIReturnError] if the general API failed.
func (a *BegetResponse[Result]) MustGetAnswer() *Answer[Result] {
	if a.IsError() {
		panic(ErrAPIReturnError)
	}

	return a.answer
}

// MustGetError returns the response of the API method call if the general API failed
// or panics with [ErrAPIReturnSuccess] if the general API did not fail.
func (a *BegetResponse[Result]) MustGetError() *Error[APIErrorCode] {
	if !a.IsError() {
		panic(ErrAPIReturnSuccess)
	}

	return a.error
}

// jsonResponse is a temporary structure to parse the general part of Beget.API response from json.
type jsonResponse[Result any] struct {
	Status    Status          `json:"status"`
	Answer    *Answer[Result] `json:"answer,omitempty"`
	ErrorText string          `json:"error_text,omitempty"`
	ErrorCode *APIErrorCode   `json:"error_code,omitempty"`
}

// UnmarshalJSON is functions for [encoding/json] to unmarshal response from json format.
func (a *BegetResponse[Result]) UnmarshalJSON(data []byte) error {
	var response jsonResponse[Result]
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	a.status = response.Status
	a.answer = response.Answer
	a.error = nil

	if response.ErrorCode != nil {
		a.error = &Error[APIErrorCode]{
			ErrorCode: *response.ErrorCode,
			ErrorText: response.ErrorText,
		}
	}

	return nil
}
