package datastore

import "errors"

var (
	ErrNoSuchKey = errors.New("datastore contains no such key")
)
