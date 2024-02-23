// Package core implements the basic functionality of requests to the Beget.API.
// The package includes the functionality for creating and sending a request,
// as well as the general part of the API response.
//
// To generate.The API has a specific response consisting of two parts.
// The response of the general API and the response of a specific method.
// More information can be founded in [BegetResponseDoc].
// In json format response of general API, it represents:
//
//	{
//	  "status": "success",
//	  "answer": {...},
//	}
//
// or:
//
//	{
//	  "status": "success",
//	  "error": "...",
//	}
//
// In json format response of a specific method, it represents:
//
//	{
//	  "status": "success",
//	  "result": {...},
//	}
//
// or:
//
//	{
//	  "status": "success",
//	  "error": [ ... ],
//	}
//
// The response of a specific method is the value of the "response" field.
//
// # Note
//
// The value of the "answer" field will be called 'the API method response',
// and the value of the "result" field will be called 'the API method result'.
//
// [BegetResponseDoc]: https://beget.com/ru/kb/api/obshhij-princzip-raboty-s-api#obrabotka-otveta.
package core

import "net/url"

// MethodName is typed of name of Beget.Api method.
type MethodName string

// APIMethod is an interface that describes the methods
// for the library core to work with the methods of the Beget.Api.
type APIMethod[Result any] interface {
	// GetHTTPMethod returns name of http method (POST, GET, etc.).
	GetHTTPMethod() string
	// GetURL returns suffix url of api method. The url is expected to contain the required query parameters.
	GetURL() *url.URL
	// Error returns any errors when generating information about the request to the method.
	Error() error
	// GetName returns name of method (to use it's in an error message).
	GetName() MethodName
}
