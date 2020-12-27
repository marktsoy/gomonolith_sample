package store

import "errors"

var (
	// ErrRecordNotFound - returned where no records found
	ErrRecordNotFound = errors.New("No records found")
)
