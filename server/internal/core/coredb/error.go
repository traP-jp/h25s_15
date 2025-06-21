package coredb

import "errors"

var ErrRecordNotFound = errors.New("not found")
var ErrNoRecordUpdated = errors.New("no record updated")
