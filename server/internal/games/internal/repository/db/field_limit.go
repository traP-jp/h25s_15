package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *Repo) InitializeFieldCardsLimit(ctx context.Context, gameID uuid.UUID, limit int) error {
	_, err := r.db.DB(ctx).ExecContext(ctx,
		"INSERT IGNORE INTO field_cards_limits (game_id, field_cards) VALUES (?, ?)",
		gameID, limit)
	if err != nil {
		return fmt.Errorf("initialize field cards limit: %w", err)
	}

	return nil
}
