package response

import (
	"strings"

	"github.com/ThCompiler/go.beget.api/pkg/beget/core/info/api"
	"github.com/ThCompiler/go.beget.api/pkg/utils/slices"
)

type MethodErrors []Error

func (m MethodErrors) Error() string {
	errorStrings := slices.Map(m, func(e Error) string { return e.Error() })

	return "method return errors: " + strings.Join(errorStrings, "; ")
}

type Error struct {
	ErrorText string        `json:"error_text"`
	ErrorCode api.ErrorCode `json:"error_code"`
}

func (e *Error) Error() string {
	return string(e.ErrorCode) + ": " + e.ErrorText
}
