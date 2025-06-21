package ws

import (
	"net/http"

	"github.com/olahol/melody"
)

type Event struct {
	m *melody.Melody
}

func New(m *melody.Melody) *Event {
	return &Event{
		m: m,
	}
}

func (e *Event) HandleRequestWithKeys(res http.ResponseWriter, req *http.Request, keys map[string]any) {
	_ = e.m.HandleRequestWithKeys(res, req, keys)
}
