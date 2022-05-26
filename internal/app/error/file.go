package errors

import (
	"errors"
	"net/http"
)

var (
	ErrFileNotFound = errors.New("file not found")

	fileErrorMapper ErrorMapper = ErrorMapper{
		ErrFileNotFound: http.StatusNotFound,
	}
)
