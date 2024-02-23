package core

// Status represents the status of request to Beget.API.
type Status string

const (
	SUCCESS = Status("success") // call is successful.
	ERROR   = Status("error")   // call is failed.
)

// Format represents the type of input and output formats for the API request and response, respectively.
type Format string

const (
	JSON = Format("json") // JSON format (supports only a json format).
)

// APIErrorCode represents the error code of the API request response.
type APIErrorCode string

const (
	AuthError        = APIErrorCode("AUTH_ERROR")        // authorization error.
	IncorrectRequest = APIErrorCode("INCORRECT_REQUEST") // an error indicating an incorrect API request.
	NoSuchMethod     = APIErrorCode("NO_SUCH_METHOD")    // the specified method does not exist.
)

// MethodErrorCode represents the error code of the method call response
type MethodErrorCode string

const (
	InvalidData = MethodErrorCode("INVALID_DATA") // validation error of the transmitted data.
	// failure to complete due to reaching any limit (for example, the limit of sites or the limit of API
	// requests is exceeded (no more than 60 requests per minute for the user))
	LimitError   = MethodErrorCode("LIMIT_ERROR")
	MethodFailed = MethodErrorCode("METHOD_FAILED") // internal error when executing the method.
)
