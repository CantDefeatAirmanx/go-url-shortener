package link_repository_errors

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnknownError   = errors.New("unknown error")
)
