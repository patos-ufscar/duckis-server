package utils

import "errors"

var (
	ErrValueTimedOut error = errors.New("value timed out")
	ErrKeyNotPresent error = errors.New("key not present in db")
)