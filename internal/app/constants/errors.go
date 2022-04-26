package constants

import "errors"

var (
	DatabaseRecordNotFound = errors.New("record not found")
	EmailAlreadyUsed       = errors.New("email already used")
)
