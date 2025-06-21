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

	// GetPlayers は、該当するgameIDのプレイヤー情報を取得する。
	GetPlayers(ctx context.Context, gameID uuid.UUID) ([]domain.Player, error)

	// CreateWaitingPlayer は、指定されたユーザー名で待機中のプレイヤーを作成する。
	// 既に待機中のプレイヤーが存在する場合は、coredb.ErrDuplicateKeyを返す。
	CreateWaitingPlayer(ctx context.Context, userName string) error
}
