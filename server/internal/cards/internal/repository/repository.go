package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/cards/internal/domain"
)

type Repo interface {
	ClearHandCards(ctx context.Context, gameID uuid.UUID, playerID int) error
	GetPlayer(ctx context.Context, gameID uuid.UUID, userName string) (domain.GamePlayer, error)
	UpdateScore(ctx context.Context, gameID uuid.UUID, playerID int, score int) error
}
