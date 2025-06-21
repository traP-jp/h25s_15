package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
)

type Repo interface {
	// GetGame は、gameIDに対応するゲーム情報を取得する。
	// 該当するゲームが存在しない場合は、coredb.ErrRecordNotFoundを返す。
	GetGame(ctx context.Context, gameID uuid.UUID) (domain.Game, error)

	// GetScores は、該当するgameIDのスコアを取得する。
	GetScores(ctx context.Context, gameID uuid.UUID) ([]domain.Player, error)
}
