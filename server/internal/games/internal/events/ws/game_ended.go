package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (e *Event) GameEnded(ctx context.Context, gameID uuid.UUID, event events.GameEndedEvent) error {
	eventJSON, err := corews.JSON(event)
	if err != nil {
		return fmt.Errorf("marshal game ended event: %w", err)
	}

	err = e.m.BroadcastBinaryFilter(eventJSON, corews.FilterGameID(gameID))
	if err != nil {
		return fmt.Errorf("broadcast game ended event: %w", err)
	}

	return nil
}
