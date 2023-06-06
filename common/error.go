package common

import "errors"

var (
	ErrInvalidResource       = errors.New("invalid resource")
	ErrAuthFailed            = errors.New("auth failed")
	ErrResourceAlreadyExists = errors.New("resource already exists")
	ErrResourceNotFound      = errors.New("resource not found")
)
