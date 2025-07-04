package repository

import (
	"context"
	"time"

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

	// DeleteWaitingPlayers は、指定されたユーザー名の待機中のプレイヤーを削除する。
	DeleteWaitingPlayers(ctx context.Context, userNames []string) error

	// GetWaitingPlayers は、待機中のプレイヤー情報を古い順に取得する。
	GetWaitingPlayers(ctx context.Context) ([]domain.WaitingPlayer, error)

	// CreateGames は、新しいゲームを作成する。
	CreateGames(ctx context.Context, gameID []uuid.UUID) error

	// CreatePlayers は、指定されたゲームIDとユーザー名でプレイヤーを作成する。
	CreatePlayers(ctx context.Context, args []CreatePlayersArg) error

	// GetActiveCards は、指定されたゲームIDのアクティブなカード情報を取得する。
	GetActiveCards(ctx context.Context, gameID uuid.UUID) ([]domain.Card, error)

	// GetGameHandLimitは、ゲームのhand cardsの制限を取得する。
	// playerID 0、playerID 1の順で、各プレイヤーのhand cardsの制限を返す。
	GetGameHandLimit(ctx context.Context, gameID uuid.UUID) ([2]int, error)

	// InitializeHandLimit は、指定されたゲームIDのプレイヤーのhand cardsの制限を初期化する。
	InitializeHandLimit(ctx context.Context, gameID uuid.UUID) error

	// InitializeFieldCardsLimit は、指定されたゲームIDのフィールドカードの制限を初期化する。
	InitializeFieldCardsLimit(ctx context.Context, gameID uuid.UUID, limit int) error

	// StartGame は、指定されたゲームIDのゲームを開始する。
	StartGame(ctx context.Context, gameID uuid.UUID, startAt time.Time) error

	// CreateTurn は、指定されたゲームIDのターンを作成する。
	CreateTurn(ctx context.Context, gameID uuid.UUID, turn int, playerID int, startAt time.Time, endAt time.Time) error

	// GetTurn は、指定されたゲームIDの最新のターン情報を取得する。
	GetTurn(ctx context.Context, gameID uuid.UUID) (domain.Turn, error)

	// EndGame は、指定されたゲームIDのゲームを終了する。
	EndGame(ctx context.Context, gameID uuid.UUID, endAt time.Time) error

	// CreateCard は、指定されたゲームIDに新しいカードをfieldで作成する。
	CreateCard(ctx context.Context, cardID uuid.UUID, gameID uuid.UUID, cardType string, value string) error

	// GetSuccessExpressions は、指定されたゲームIDの成功した式を取得する。
	GetSuccessExpressions(ctx context.Context, gameID uuid.UUID) ([]domain.Expression, error)

	// GetRanking は、ランキング情報を取得する。
	// ランキングは、勝利数が多い順に並べられ、同じ勝利数の場合は合計得点が多い方が上位になる。
	GetRanking(ctx context.Context, limit int) ([]domain.RankingItem, error)

	GetUsersCount(ctx context.Context) (int, error)
}

type CreatePlayersArg struct {
	GameID    uuid.UUID
	UserName0 string
	UserName1 string
}
