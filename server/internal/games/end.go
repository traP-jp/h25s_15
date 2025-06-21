package games

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (h *Handler) EndGame(ctx context.Context, gameID uuid.UUID, endAt time.Time) error {
	err := h.db.Transaction(ctx, func(ctx context.Context) error {
		err := h.repo.EndGame(ctx, gameID, endAt)
		if err != nil {
			return fmt.Errorf("end game: %w", err)
		}

		players, err := h.repo.GetPlayers(ctx, gameID)
		if err != nil {
			return fmt.Errorf("get players: %w", err)
		}

		var score0, score1 int
		for _, player := range players {
			switch player.PlayerID {
			case 0:
				score0 = player.Score
			case 1:
				score1 = player.Score
			default:
				return fmt.Errorf("unexpected player ID: %d", player.PlayerID)
			}
		}

		err = h.events.GameEnded(ctx, gameID, events.GameEndedEvent{
			Type:    "gameEnded",
			Player0: score0,
			Player1: score1,
		})
		if err != nil {
			return fmt.Errorf("send game ended event: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("end game transaction: %w", err)
	}

	return nil
}
