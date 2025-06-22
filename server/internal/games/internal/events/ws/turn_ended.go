package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (e *Event) TurnEnded(ctx context.Context, gameID uuid.UUID, event events.TurnEndedEvent) error {
	eventJSON, err := corews.JSON(event)
	if err != nil {
		return fmt.Errorf("marshal turn ended event: %w", err)
	}

	err = e.m.BroadcastFilter(eventJSON, corews.FilterGameID(gameID))
	if err != nil {
		return fmt.Errorf("broadcast turn ended event: %w", err)
	}

	return nil
}
