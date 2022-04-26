package errors

import (
	"errors"
	"net/http"
)

const ()

var (
	ErrUserEmailNotFound    = errors.New("user with that email not found")
	ErrEmailAlreadyExist    = errors.New("email already exist")
	ErrWrongPassword        = errors.New("wrong password. please check again your password")
	ErrPasswordDoesNotMatch = errors.New("password confirmation does not match")
	ErrMissingAuthorization = errors.New("missing authorization header")
	ErrWrongAuthorization   = errors.New("wrong authorization header format")
	ErrDecodingJWT          = errors.New("something wrong in decoding JWT")
	ErrRoleUnauthorized     = errors.New("role unauthorized")

	userErrorMapper ErrorMapper = ErrorMapper{
		ErrUserEmailNotFound:    http.StatusNotFound,
		ErrEmailAlreadyExist:    http.StatusBadRequest,
		ErrWrongPassword:        http.StatusBadRequest,
		ErrPasswordDoesNotMatch: http.StatusBadRequest,
		ErrMissingAuthorization: http.StatusUnauthorized,
		ErrWrongAuthorization:   http.StatusUnauthorized,
		ErrRoleUnauthorized:     http.StatusUnauthorized,
	}
)
