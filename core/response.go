package core

import (
	"encoding/json"
)

// Error is the type to represent an API method call general's errors.
type Error[Err MethodErrorCode | APIErrorCode] struct {
	ErrorText string `json:"error_text"`
	ErrorCode Err    `json:"error_code"`
}

// Error represents Error as a string.
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
func (br *BegetResponse[Result]) Get() (*Answer[Result], error) {
	if br.error == nil {
		return br.answer, nil
	}

	return br.answer, br.error
}

// Status returns the response status of the general API.
func (br *BegetResponse[Result]) Status() Status {
	return br.status
}

// IsError reports whether the general API returns some error.
func (br *BegetResponse[Result]) IsError() bool {
	return br.error != nil
}

// GetAnswer returns the response of the API method call if the general API did not fail
// or [ErrAPIReturnError] if the general API failed.
func (br *BegetResponse[Result]) GetAnswer() (*Answer[Result], error) {
	if br.IsError() {
		return nil, ErrAPIReturnError
	}

	return br.answer, nil
}

// GetError returns the response of the API method call if the general API failed
// or [ErrAPIReturnSuccess] if the general API did not fail.
func (br *BegetResponse[Result]) GetError() (*Error[APIErrorCode], error) {
	if !br.IsError() {
		return nil, ErrAPIReturnSuccess
	}

	return br.error, nil
}

// MustGetAnswer returns the response of the API method call if the general API did not fail
// or panics with [ErrAPIReturnError] if the general API failed.
func (br *BegetResponse[Result]) MustGetAnswer() *Answer[Result] {
	if br.IsError() {
		panic(ErrAPIReturnError)
	}

	return br.answer
}

// MustGetError returns the response of the API method call if the general API failed
// or panics with [ErrAPIReturnSuccess] if the general API did not fail.
func (br *BegetResponse[Result]) MustGetError() *Error[APIErrorCode] {
	if !br.IsError() {
		panic(ErrAPIReturnSuccess)
	}

	return br.error
}

// jsonResponse is a temporary structure to parse the general part of Beget.API response from json.
type jsonResponse[Result any] struct {
	Status    Status          `json:"status"`
	Answer    *Answer[Result] `json:"answer,omitempty"`
	ErrorText string          `json:"error_text,omitempty"`
	ErrorCode *APIErrorCode   `json:"error_code,omitempty"`
}

// UnmarshalJSON is functions for [encoding/json] to unmarshal response from json format.
func (br *BegetResponse[Result]) UnmarshalJSON(data []byte) error {
	var response jsonResponse[Result]
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	br.status = response.Status
	br.answer = response.Answer
	br.error = nil

	if response.ErrorCode != nil {
		br.error = &Error[APIErrorCode]{
			ErrorCode: *response.ErrorCode,
			ErrorText: response.ErrorText,
		}
	}

	return nil
}
