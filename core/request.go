package core

import (
	"encoding/json"
	"net/http"
)

// BegetRequest is a prepared request to Beget,API.
// It contains prepared [http.Request] and [http.Client].
// The Result is a generic parameter that allows the query executor to parse
// the query result with the previously set result type.
type BegetRequest[Result any] struct {
	req    *http.Request
	client *http.Client
	_      Result
}

// NewBegetRequest creates a new [BegetRequest].
func NewBegetRequest[Result any](req *http.Request, client *http.Client) *BegetRequest[Result] {
	if client == nil {
		client = http.DefaultClient
	}
	return &BegetRequest[Result]{
		req:    req,
		client: client,
	}
}

// Do executes prepared request and parse response to [BegetResponse].
// It may return an http request error or a json parsing error.
func (r *BegetRequest[Result]) Do() (*BegetResponse[Result], error) {
	resp, err := r.client.Do(r.req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	decoder := json.NewDecoder(resp.Body)

	var begetResponse BegetResponse[Result]
	if err := decoder.Decode(&begetResponse); err != nil {
		return nil, err
	}

	return &begetResponse, nil
}
