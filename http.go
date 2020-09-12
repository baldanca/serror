package serror

import "net/http"

type (
	// HTTPError model
	HTTPError struct {
		*Error
		Status     string
		StatusCode int
		*Message
	}

	// Message type
	Message string
)

// HTTPCustom is to custom error
func (e *Error) HTTPCustom(statusCode int, m *Message) *HTTPError {
	return &HTTPError{
		e,
		http.StatusText(statusCode),
		statusCode,
		m,
	}
}

// HTTPInternal is to every internal server error
func (e *Error) HTTPInternal(m *Message) *HTTPError {
	return e.HTTPCustom(http.StatusInternalServerError, m)
}

// HTTPbad is to every bad request error
func (e *Error) HTTPbad(m *Message) *HTTPError {
	return e.HTTPCustom(http.StatusBadRequest, m)
}

// HTTPUnauthorized is to every Unauthorized request
func (e *Error) HTTPUnauthorized(m *Message) *HTTPError {
	return e.HTTPCustom(http.StatusUnauthorized, m)
}

// HTTPGeneric is to every generic error
func (e *Error) HTTPGeneric() *HTTPError {
	return e.HTTPInternal(&Default)
}
