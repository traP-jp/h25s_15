package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/events"
)

func (e *Event) SubmissionSucceeded(ctx context.Context, gameID uuid.UUID, event events.SubmissionSucceededEvent) error {
	eventJSON, err := corews.JSON(event)
	if err != nil {
		return fmt.Errorf("marshal event to JSON: %w", err)
	}

	err = e.m.BroadcastFilter(eventJSON, corews.FilterGameID(gameID))
	if err != nil {
		return fmt.Errorf("broadcast event: %w", err)
	}

	return nil
}
