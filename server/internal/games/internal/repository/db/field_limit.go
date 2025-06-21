package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

const initialFieldCardsLimit int = 4

func (r *Repo) InitializeFieldCardsLimit(ctx context.Context, gameID uuid.UUID) error {
	_, err := r.db.DB(ctx).ExecContext(ctx,
		"INSERT IGNORE INTO field_cards_limits (game_id, field_cards) VALUES (?, ?)",
		gameID, initialFieldCardsLimit)
	if err != nil {
		return fmt.Errorf("initialize field cards limit: %w", err)
	}

	return nil
}
