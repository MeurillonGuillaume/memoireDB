package datastore

import "errors"

var (
	ErrNoSuchKey        = errors.New("datastore contains no such key(s)")
	ErrPrefixWhitespace = errors.New("query prefix contains nothing but space, add a prefix or use an empty prefix")
)
