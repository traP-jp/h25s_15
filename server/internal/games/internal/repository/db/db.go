package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
	"github.com/traP-jp/h25s_15/internal/games/internal/repository"
)

type Repo struct {
	db *coredb.DB
}

func New(db *coredb.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetGame(ctx context.Context, gameID uuid.UUID) (domain.Game, error) {
	var game Game
	err := r.db.DB(ctx).GetContext(ctx, &game, "SELECT * FROM games WHERE id = ?", gameID)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Game{}, coredb.ErrRecordNotFound
	}
	if err != nil {
		return domain.Game{}, fmt.Errorf("get game: %w", err)
	}

	var startedAt *time.Time
	if game.StartedAt.Valid {
		startedAt = &game.StartedAt.Time
	}
	var endedAt *time.Time
	if game.EndedAt.Valid {
		endedAt = &game.EndedAt.Time
	}
	result := domain.Game{
		ID:        game.ID,
		Status:    domain.GameStatus(game.Status),
		StartedAt: startedAt,
		EndedAt:   endedAt,
		CreatedAt: game.CreatedAt,
	}

	return result, nil
}

func (r *Repo) GetPlayers(ctx context.Context, gameID uuid.UUID) ([]domain.Player, error) {
	var players []Player
	err := r.db.DB(ctx).SelectContext(ctx, &players, "SELECT * FROM game_players WHERE game_id = ?", gameID)
	if err != nil {
		return nil, fmt.Errorf("get players: %w", err)
	}

	result := make([]domain.Player, 0, len(players))
	for _, player := range players {
		result = append(result, domain.Player{
			GameID:   player.GameID,
			PlayerID: player.PlayerID,
			UserName: player.UserName,
			Score:    player.Score,
		})
	}

	return result, nil
}

func (r *Repo) CreateWaitingPlayer(ctx context.Context, userName string) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "INSERT INTO waiting_players (user_name, created_at) VALUES (?, ?)", userName, time.Now())
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 { // unique key 制約違反
			return coredb.ErrDuplicateKey
		}
		return fmt.Errorf("create waiting player: %w", err)
	}

	return nil
}

func (r *Repo) DeleteWaitingPlayers(ctx context.Context, userNames []string) error {
	q, args, err := sqlx.In("DELETE FROM waiting_players WHERE user_name IN (?)", userNames)
	if err != nil {
		return fmt.Errorf("prepare delete waiting players query: %w", err)
	}

	_, err = r.db.DB(ctx).ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("delete waiting players: %w", err)
	}

	return nil
}

func (r *Repo) GetWaitingPlayers(ctx context.Context) ([]domain.WaitingPlayer, error) {
	var waitingPlayers []WaitingPlayer
	err := r.db.DB(ctx).SelectContext(ctx, &waitingPlayers,
		"SELECT * FROM waiting_players ORDER BY created_at ASC")
	if err != nil {
		return nil, fmt.Errorf("get waiting players: %w", err)
	}

	result := make([]domain.WaitingPlayer, 0, len(waitingPlayers))
	for _, player := range waitingPlayers {
		result = append(result, domain.WaitingPlayer{
			UserName:  player.UserName,
			CreatedAt: player.CreatedAt,
		})
	}

	return result, nil
}

func (r *Repo) CreateGames(ctx context.Context, gameID []uuid.UUID) error {
	games := make([]Game, 0, len(gameID))
	for _, id := range gameID {
		games = append(games, Game{
			ID:     id,
			Status: string(domain.GameStatusWaiting),
		})
	}
	_, err := r.db.DB(ctx).NamedExecContext(ctx,
		"INSERT INTO games (id, status) VALUES (:id, :status)",
		games,
	)
	if err != nil {
		return fmt.Errorf("create games: %w", err)
	}

	return nil
}

func (r *Repo) CreatePlayers(ctx context.Context, args []repository.CreatePlayersArg) error {
	players := make([]Player, 0, len(args))
	for _, arg := range args {
		players = append(players, Player{
			GameID:   arg.GameID,
			UserName: arg.UserName0,
			PlayerID: 0,
			Score:    0, // 初期スコアは0
		})
		players = append(players, Player{
			GameID:   arg.GameID,
			UserName: arg.UserName1,
			PlayerID: 1,
			Score:    0, // 初期スコアは0
		})
	}

	_, err := r.db.DB(ctx).NamedExecContext(ctx,
		"INSERT INTO game_players (game_id, player_id, user_name, score) VALUES (:game_id, :player_id, :user_name, :score)",
		players,
	)
	if err != nil {
		return fmt.Errorf("create players: %w", err)
	}

	return nil

}

func (r *Repo) GetActiveCards(ctx context.Context, gameID uuid.UUID) ([]domain.Card, error) {
	var cards []Card
	err := r.db.DB(ctx).SelectContext(ctx, &cards, "SELECT * FROM cards WHERE game_id = ? AND location IN (?, ?)",
		gameID, domain.CardLocationHand, domain.CardLocationField)
	if err != nil {
		return nil, fmt.Errorf("get active cards: %w", err)
	}

	domainCards := make([]domain.Card, 0, len(cards))
	for _, card := range cards {
		var playerID *int
		if card.OwnerPlayerID.Valid {
			id := int(card.OwnerPlayerID.Int16)
			playerID = &id
		}
		domainCards = append(domainCards, domain.Card{
			ID:            card.ID,
			Type:          domain.CardType(card.Type),
			Value:         card.Value,
			Location:      domain.CardLocation(card.Location),
			OwnerPlayerID: playerID,
		})
	}
	return domainCards, nil
}

func (r *Repo) GetGameHandLimit(ctx context.Context, gameID uuid.UUID) ([2]int, error) {
	var handLimits []HandLimit
	err := r.db.DB(ctx).SelectContext(ctx, &handLimits,
		"SELECT * FROM hand_cards_limits WHERE game_id = ? ORDER BY player_id ASC", gameID)
	if err != nil {
		return [2]int{}, fmt.Errorf("get game hand limit: %w", err)
	}
	if len(handLimits) < 2 {
		return [2]int{}, fmt.Errorf("get game hand limit: not enough data")
	}

	return [2]int{handLimits[0].HandCardsLimit, handLimits[1].HandCardsLimit}, nil
}

const initialHandLimit int = 10 // 初期のhand cardsの制限

func (r *Repo) InitializeHandLimit(ctx context.Context, gameID uuid.UUID) error {
	_, err := r.db.DB(ctx).ExecContext(ctx,
		"INSERT INTO hand_cards_limits (game_id, player_id, hand_cards) VALUES (?, 0, ?), (?, 1, ?)",
		gameID, initialHandLimit, gameID, initialHandLimit)
	if err != nil {
		return fmt.Errorf("initialize hand limit: %w", err)
	}

	return nil
}

func (r *Repo) StartGame(ctx context.Context, gameID uuid.UUID, startAt time.Time) error {
	_, err := r.db.DB(ctx).ExecContext(ctx,
		"UPDATE games SET status = ?, started_at = ? WHERE id = ?",
		domain.GameStatusRunning, sql.NullTime{Time: startAt, Valid: true}, gameID,
	)
	if err != nil {
		return fmt.Errorf("start game: %w", err)
	}

	return nil
}

func (r *Repo) CreateTurn(ctx context.Context, gameID uuid.UUID, turn int, playerID int, startAt time.Time, endAt time.Time) error {
	_, err := r.db.DB(ctx).ExecContext(ctx,
		"INSERT INTO turns (game_id, turn_number, player_id, start_at, end_at) VALUES (?, ?, ?, ?, ?)",
		gameID, turn, playerID, startAt, endAt,
	)
	if err != nil {
		return fmt.Errorf("create turn: %w", err)
	}

	return nil
}

func (r *Repo) GetTurn(ctx context.Context, gameID uuid.UUID) (domain.Turn, error) {
	var turn Turn
	err := r.db.DB(ctx).
		GetContext(ctx, &turn,
			"SELECT * FROM turns WHERE game_id = ? ORDER BY turn_number DESC LIMIT 1",
			gameID)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Turn{}, coredb.ErrRecordNotFound
	}
	if err != nil {
		return domain.Turn{}, fmt.Errorf("get turn: %w", err)
	}

	return domain.Turn{
		GameID:     turn.GameID,
		PlayerID:   turn.PlayerID,
		TurnNumber: turn.TurnNumber,
		StartAt:    turn.StartAt,
		EndAt:      turn.EndAt,
	}, nil
}

func (r *Repo) EndGame(ctx context.Context, gameID uuid.UUID, endAt time.Time) error {
	result, err := r.db.DB(ctx).ExecContext(ctx,
		"UPDATE games SET status = ?, ended_at = ? WHERE id = ?",
		domain.GameStatusFinished, sql.NullTime{Time: endAt, Valid: true}, gameID,
	)
	if err != nil {
		return fmt.Errorf("end game: %w", err)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("get rows affected: %w", err)
	} else if rowsAffected == 0 {
		return coredb.ErrRecordNotFound
	}

	return nil
}
