package expressions

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/events"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/events/ws"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/repository"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/repository/db"
)

type Handler struct {
	db     *coredb.DB
	repo   repository.Repo
	events events.Event
	parser *participle.Parser[Expr]
}

func New(d *coredb.DB, m *melody.Melody) (*Handler, error) {
	repo := db.New(d)
	events := ws.New(m)
	parser, err := Parser()
	if err != nil {
		return nil, fmt.Errorf("failed to create expression parser: %w", err)
	}

	return &Handler{
		db:     d,
		events: events,
		repo:   repo,
		parser: parser,
	}, nil
}
