package session

import "errors"

const (
	NAME = "session"
)

var (
	errInvalidSessionConnection = errors.New("Invalid session connection.")
)
