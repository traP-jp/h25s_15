package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
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
	err := r.db.DB(ctx).SelectContext(ctx, &players, "SELECT * FROM players WHERE game_id = ?", gameID)
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
