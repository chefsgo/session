package session

import "errors"

const (
	NAME = "SESSION"
)

var (
	errInvalidSessionConnection = errors.New("Invalid session connection.")
)
