package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/domain"
)

type Repo interface {
	// GetCards retrieves cards by their IDs.
	GetCards(ctx context.Context, gameID uuid.UUID, playerID int, cardIDs []uuid.UUID) ([]domain.Card, error)

	// GetPlayers retrieves players for a given game.
	GetPlayers(ctx context.Context, gameID uuid.UUID) ([]domain.GamePlayer, error)

	// UseCards marks the specified cards as used by a player in a game.
	UseCards(ctx context.Context, gameID uuid.UUID, playerID int, cardIDs []uuid.UUID) error

	// UpdatePlayerScore updates the score of a player in a game.
	UpdatePlayerScore(ctx context.Context, gameID uuid.UUID, playerID int, scoreDiff int) error

	// CreateExpression logs the creation of an expression by a player in a game.
	CreateExpression(ctx context.Context, id uuid.UUID, gameID uuid.UUID, playerID int, expression string, value string, point int, success bool) error
}
