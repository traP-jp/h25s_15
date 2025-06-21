package ws

import "github.com/olahol/melody"

type Event struct {
	m *melody.Melody
}

func New(m *melody.Melody) *Event {
	return &Event{
		m: m,
	}
}
