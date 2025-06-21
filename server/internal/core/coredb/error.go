package coredb

import "errors"

var (
	ErrRecordNotFound  = errors.New("not found")
	ErrNoRecordUpdated = errors.New("no record updated")
	ErrDuplicateKey    = errors.New("duplicate key")
)
