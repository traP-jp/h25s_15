package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/cards/internal/domain"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
)

type Repo struct {
	db *coredb.DB
}

func New(db *coredb.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) ClearHandCards(ctx context.Context, gameID uuid.UUID, playerID int) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE cards SET location = 'used' WHERE game_id = ? and owner_player_id = ?", gameID, playerID)
	if err != nil {
		return fmt.Errorf("update location: %w", err)
	}
	return nil
}

func (r *Repo) GetPlayer(ctx context.Context, gameID uuid.UUID, userName string) (domain.GamePlayer, error) {
	var player Player
	err := r.db.DB(ctx).GetContext(ctx, &player, "SELECT * FROM game_players WHERE game_id = ? and user_name = ?", gameID, userName)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.GamePlayer{}, coredb.ErrRecordNotFound
	}
	if err != nil {
		return domain.GamePlayer{}, fmt.Errorf("get player: %w", err)
	}
	return domain.GamePlayer(player), nil
}

func (r *Repo) UpdateScore(ctx context.Context, gameID uuid.UUID, playerID int, diff int) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE game_players SET score = score + (?) WHERE game_id = ? and player_id = ?",
		diff, gameID, playerID)
	if err != nil {
		return fmt.Errorf("update score: %w", err)
	}
	return nil
}

func (r *Repo) PickFieldCards(ctx context.Context, gameID uuid.UUID, playerID int, requestCard uuid.UUID) error {
	result, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE cards SET location = 'hand', owner_player_id = ? WHERE game_id = ? and id = ?",
		playerID, gameID, requestCard)
	if err != nil {
		return fmt.Errorf("update location: %w", err)
	}
	ok, _ := result.RowsAffected()
	if ok != 1 {
		return coredb.ErrNoRecordUpdated
	}
	return nil
}
