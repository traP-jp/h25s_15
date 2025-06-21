package cards

import (
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/cards/internal/events"
	"github.com/traP-jp/h25s_15/internal/cards/internal/events/ws"
	"github.com/traP-jp/h25s_15/internal/cards/internal/repository"
	"github.com/traP-jp/h25s_15/internal/cards/internal/repository/db"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
)

type Handler struct {
	db     *coredb.DB
	repo   repository.Repo
	events events.Event
}

func New(d *coredb.DB, m *melody.Melody) *Handler {
	repo := db.New(d)
	events := ws.New(m, repo)

	return &Handler{
		db:     d,
		events: events,
		repo:   repo,
	}
}
