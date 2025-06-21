package events

import (
	"net/http"
)

type Event interface {
	HandleRequestWithKeys(res http.ResponseWriter, req *http.Request, keys map[string]any)
}
