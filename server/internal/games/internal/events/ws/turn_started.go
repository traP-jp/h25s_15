package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (e *Event) TurnStarted(ctx context.Context, gameID uuid.UUID, event events.TurnStartedEvent) error {
	eventJSON, err := corews.JSON(event)
	if err != nil {
		return fmt.Errorf("marshal turn started event: %w", err)
	}

	err = e.m.BroadcastFilter(eventJSON, corews.FilterGameID(gameID))
	if err != nil {
		return fmt.Errorf("broadcast turn started event: %w", err)
	}

	return nil
}
