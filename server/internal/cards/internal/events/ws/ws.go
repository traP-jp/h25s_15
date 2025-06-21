package ws

import (
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/cards/internal/repository"
)

type Event struct {
	m    *melody.Melody
	repo repository.Repo
}

func New(m *melody.Melody, repo repository.Repo) *Event {
	return &Event{
		m:    m,
		repo: repo,
	}
}
