package derpigo

import "net/http"

/*
Error is a combination of a Go error and a Derpibooru request ID to
help with debugging failed API calls with the Derpibooru staff.
*/
type Error struct {
	Underlying error
	RequestID  string
}

// NewError wraps an error with the X-Request-Id.
func NewError(underlying error, resp *http.Response) *Error {
	return &Error{
		Underlying: underlying,
		RequestID:  resp.Header.Get("X-Request-Id"),
	}
}

// Error satisfies the error interface.
func (e *Error) Error() string {
	return e.Underlying.Error() + " Request: " + e.RequestID
}
