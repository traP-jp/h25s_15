package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/domain"
)

type Repo struct {
	db *coredb.DB
}

func New(db *coredb.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetCards(ctx context.Context, gameID uuid.UUID, playerID int, cardIDs []uuid.UUID) ([]domain.Card, error) {
	cards := make([]Card, 0, len(cardIDs))
	query, args, err := sqlx.In(`
		SELECT * FROM cards
		WHERE game_id = ? AND owner_player_id = ? AND id IN (?)
	`, gameID, playerID, cardIDs)
	if err != nil {
		return nil, fmt.Errorf("prepare query for cards: %w", err)
	}

	err = r.db.DB(ctx).SelectContext(ctx, &cards, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select cards: %w", err)
	}

	result := make([]domain.Card, 0, len(cards))
	for _, card := range cards {
		var playerID *int
		if card.OwnerPlayerID.Valid {
			p := int(card.OwnerPlayerID.Int16)
			playerID = &p
		}
		result = append(result, domain.Card{
			ID:   card.ID,
			Type: domain.CardType(card.Type),

			Value:         card.Value,
			OwnerPlayerID: playerID,
			Location:      domain.CardLocation(card.Location),
		})
	}

	return result, nil
}

func (r *Repo) GetPlayers(ctx context.Context, gameID uuid.UUID) ([]domain.GamePlayer, error) {
	players := make([]Player, 0)
	query := `SELECT * FROM game_players WHERE game_id = ?`
	err := r.db.DB(ctx).SelectContext(ctx, &players, query, gameID)
	if err != nil {
		return nil, fmt.Errorf("select players: %w", err)
	}

	result := make([]domain.GamePlayer, 0, len(players))
	for _, player := range players {
		result = append(result, domain.GamePlayer{
			GameID:   player.GameID,
			PlayerID: player.PlayerID,
			UserName: player.UserName,
			Score:    player.Score,
		})
	}

	return result, nil
}

func (r *Repo) UseCards(ctx context.Context, gameID uuid.UUID, playerID int, cardIDs []uuid.UUID) error {
	query, args, err := sqlx.In(`
		UPDATE cards
		SET location = 'used'
		WHERE game_id = ? AND owner_player_id = ? AND id IN (?)
	`, gameID, playerID, cardIDs)
	if err != nil {
		return fmt.Errorf("prepare query for using cards: %w", err)
	}

	result, err := r.db.DB(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec update cards: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return coredb.ErrNoRecordUpdated
	}

	return nil
}

func (r *Repo) UpdatePlayerScore(ctx context.Context, gameID uuid.UUID, playerID int, scoreDiff int) error {
	query := `
		UPDATE game_players
		SET score = score + ?
		WHERE game_id = ? AND player_id = ?
	`
	result, err := r.db.DB(ctx).ExecContext(ctx, query, scoreDiff, gameID, playerID)
	if err != nil {
		return fmt.Errorf("exec update player score: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return coredb.ErrNoRecordUpdated
	}

	return nil
}

func (r *Repo) CreateExpression(ctx context.Context, id uuid.UUID, gameID uuid.UUID, playerID int, expression string, value string, point int, success bool) error {
	query := `
		INSERT INTO expressions (id, game_id, player_id, expression, value, point, success, submitted_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.DB(ctx).ExecContext(ctx, query,
		id, gameID, playerID, expression, value, point, success, time.Now())
	if err != nil {
		return fmt.Errorf("exec insert expression: %w", err)
	}

	return nil
}
