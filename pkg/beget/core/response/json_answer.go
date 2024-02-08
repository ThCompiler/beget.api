package response

import (
	"github.com/ThCompiler/go.beget.api/pkg/beget/core/info"
)

type jsonAnswer[Result any] struct {
	Status info.Status  `json:"status"`
	Errors MethodErrors `json:"errors,omitempty"`
	Result *Result      `json:"result,omitempty"`
}
