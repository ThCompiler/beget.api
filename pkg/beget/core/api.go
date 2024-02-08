package core

import "net/url"

type MethodName string

type APIMethod[Result any] interface {
	GetHTTPMethod() string
	GetURL() *url.URL
	Error() error
	GetName() MethodName
}
