package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (e *Event) ScoreUpdated(ctx context.Context, gameID uuid.UUID, event events.ScoreUpdatedEvent) error {
	eventJSON, err := corews.JSON(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event to JSON: %w", err)
	}

	err = e.m.BroadcastBinaryFilter(eventJSON, corews.FilterGameID(gameID))
	if err != nil {
		return fmt.Errorf("broadcast score updated event: %w", err)
	}

	return nil
}
