package games

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (h *Handler) StartGame(ctx context.Context, gameID uuid.UUID, startAt time.Time) error {
	<-time.After(time.Until(startAt))

	err := h.db.Transaction(ctx, func(ctx context.Context) error {
		err := h.repo.StartGame(ctx, gameID, startAt)
		if err != nil {
			return fmt.Errorf("start game: %w", err)
		}

		err = h.events.GameStarted(ctx, gameID, events.GameStartedEvent{
			Type:            "gameStarted",
			CurrentPlayerID: 0,
			Turn:            1,
		})
		if err != nil {
			return fmt.Errorf("send game started event: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("start game transaction: %w", err)
	}

	return nil
}
