package models

import "errors"

var (
	ErrValueTimedOut error = errors.New("value timed out")
)