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
	PickFieldCards(ctx context.Context, gameID uuid.UUID, playerID int, wantedcard uuid.UUID) error
	ReplenishFieldCards(ctx context.Context, gameID uuid.UUID, number int) error
	// GetActiveCardsは、ゲームのhand cardsとfield cardsを取得する。
	GetActiveCards(ctx context.Context, gameID uuid.UUID) ([]domain.Card, error)
	// GetGameHandLimitは、ゲームのhand cardsの制限を取得する。
	// playerID 0、playerID 1の順で、各プレイヤーのhand cardsの制限を返す。
	GetGameHandLimit(ctx context.Context, gameID uuid.UUID) ([2]int, error)
}
