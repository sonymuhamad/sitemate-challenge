package httpserver

import (
	"net/http"
)

type HTTPError interface {
	ToHTTPError() (int, []HTTPResponseError)
}

type HTTPResponseError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type UnprocessableError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func (e UnprocessableError) Error() string {
	return e.Message
}

func (e UnprocessableError) ToHTTPError() (int, []HTTPResponseError) {
	return http.StatusUnprocessableEntity, []HTTPResponseError{
		{
			Field:   e.Field,
			Message: e.Message,
		},
	}
}
