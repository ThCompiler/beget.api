package request

import (
	"encoding/json"
	"net/http"

	"github.com/ThCompiler/go.beget.api/pkg/beget/core/response"
)

type BegetRequest[Result any] struct {
	req    *http.Request
	client *http.Client
	_      Result
}

func NewBegetRequest[Result any](req *http.Request) *BegetRequest[Result] {
	return &BegetRequest[Result]{
		req:    req,
		client: http.DefaultClient,
	}
}

func (r *BegetRequest[Result]) Do() (*response.BegetResponse[Result], error) {
	resp, err := r.client.Do(r.req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	decoder := json.NewDecoder(resp.Body)

	var begetResponse response.BegetResponse[Result]
	if err := decoder.Decode(&begetResponse); err != nil {
		return nil, err
	}

	return &begetResponse, nil
}
