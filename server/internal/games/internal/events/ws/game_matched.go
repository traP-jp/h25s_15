package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (e *Event) GameMatched(ctx context.Context, userNames [2]string, gameID uuid.UUID) error {
	event0 := events.GameMatchedEvent{
		GameID:   gameID,
		PlayerID: 0,
	}
	event0JSON, err := corews.JSON(event0)
	if err != nil {
		return fmt.Errorf("failed to marshal event0: %w", err)
	}

	event1 := events.GameMatchedEvent{
		GameID:   gameID,
		PlayerID: 1,
	}
	event1JSON, err := corews.JSON(event1)
	if err != nil {
		return fmt.Errorf("failed to marshal event1: %w", err)
	}

	if err = e.m.BroadcastFilter(event0JSON, corews.FilterUserName(userNames[0])); err != nil {
		return fmt.Errorf("broadcast event0: %w", err)
	}

	if err = e.m.BroadcastFilter(event1JSON, corews.FilterUserName(userNames[1])); err != nil {
		return fmt.Errorf("broadcast event1: %w", err)
	}

	return nil
}
