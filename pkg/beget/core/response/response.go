package response

import (
	"encoding/json"

	"github.com/ThCompiler/go.beget.api/pkg/beget/core/info"
)

type BegetResponse[Result any] struct {
	status info.Status
	answer *Answer[Result]
	error  *Error
}

func (a *BegetResponse[Result]) Get() (*Answer[Result], error) {
	return a.answer, a.error
}

func (a *BegetResponse[Result]) Status() info.Status {
	return a.status
}

func (a *BegetResponse[Result]) IsError() bool {
	return a.error != nil
}

func (a *BegetResponse[Result]) GetResult() (*Answer[Result], error) {
	if a.IsError() {
		return nil, ErrAPIReturnError
	}

	return a.answer, nil
}

func (a *BegetResponse[Result]) GetError() (*Error, error) {
	if !a.IsError() {
		return nil, ErrAPIReturnSuccess
	}

	return a.error, nil
}

func (a *BegetResponse[Result]) MustGetResult() *Answer[Result] {
	if a.IsError() {
		panic(ErrAPIReturnError)
	}

	return a.answer
}

func (a *BegetResponse[Result]) MustGetError() *Error {
	if !a.IsError() {
		panic(ErrAPIReturnSuccess)
	}

	return a.error
}

func (a *BegetResponse[Result]) UnmarshalJSON(data []byte) error {
	var response jsonResponse[Result]
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	a.status = response.Status
	a.answer = response.Answer
	a.error = nil

	if response.ErrorCode != nil {
		a.error = &Error{
			ErrorCode: *response.ErrorCode,
			ErrorText: response.ErrorText,
		}
	}

	return nil
}
