package response

import (
	"encoding/json"

	"github.com/ThCompiler/go.beget.api/pkg/beget/core/info"
)

type Answer[Result any] struct {
	status info.Status
	errors MethodErrors
	result *Result
}

func (a *Answer[Result]) Get() (*Result, error) {
	if a.errors == nil {
		return a.result, nil
	}
	return a.result, a.errors
}

func (a *Answer[Result]) Status() info.Status {
	return a.status
}

func (a *Answer[Result]) IsError() bool {
	return a.errors != nil
}

func (a *Answer[Result]) GetResult() (*Result, error) {
	if a.IsError() {
		return nil, ErrAPIMethodReturnError
	}

	return a.result, nil
}

func (a *Answer[Result]) GetError() (MethodErrors, error) {
	if !a.IsError() {
		return nil, ErrAPIMethodReturnSuccess
	}

	return a.errors, nil
}

func (a *Answer[Result]) MustGetResult() *Result {
	if a.IsError() {
		panic(ErrAPIMethodReturnError)
	}

	return a.result
}

func (a *Answer[Result]) MustGetError() MethodErrors {
	if !a.IsError() {
		panic(ErrAPIMethodReturnSuccess)
	}

	return a.errors
}

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
