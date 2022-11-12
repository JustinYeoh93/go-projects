package errors

import "fmt"

type HttpError struct {
	Type    string
	Message string
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func DataAccessLayerError(e string) HttpError {
	return HttpError{
		Type:    "DataAccessLayerError",
		Message: e,
	}
}
