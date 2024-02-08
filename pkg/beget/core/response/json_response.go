package response

import (
	"github.com/ThCompiler/go.beget.api/pkg/beget/core/info"
	"github.com/ThCompiler/go.beget.api/pkg/beget/core/info/api"
)

type jsonResponse[Result any] struct {
	Status    info.Status     `json:"status"`
	Answer    *Answer[Result] `json:"answer,omitempty"`
	ErrorText string          `json:"error_text,omitempty"`
	ErrorCode *api.ErrorCode  `json:"error_code,omitempty"`
}
