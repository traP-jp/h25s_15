package items

import (
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/items/events"
	"github.com/traP-jp/h25s_15/internal/items/events/ws"
	"github.com/traP-jp/h25s_15/internal/items/repository"
	"github.com/traP-jp/h25s_15/internal/items/repository/db"
)

type Handler struct {
	db     *coredb.DB
	repo   repository.Repo
	events events.Event
}

func New(d *coredb.DB, m *melody.Melody) *Handler {
	repo := db.New(d)
	events := ws.New(m)

	return &Handler{
		db:     d,
		events: events,
		repo:   repo,
	}
}
