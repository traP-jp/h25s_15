package games

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

const turnsCount = 20
const turnTimeLimit = 15 * time.Second

func (h *Handler) RunTurns(ctx context.Context, gameID uuid.UUID) error {
	for turn := 1; turn <= turnsCount; turn++ {
		playerID := (turn + 1) % 2
		err := h.turn(ctx, gameID, turn, playerID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) turn(ctx context.Context, gameID uuid.UUID, turn int, playerID int) error {
	startAt := time.Now()
	endAt := startAt.Add(turnTimeLimit)
	err := h.db.Transaction(ctx, func(ctx context.Context) error {
		err := h.repo.CreateTurn(ctx, gameID, turn, playerID, startAt, endAt)
		if err != nil {
			return fmt.Errorf("create turn: %w", err)
		}

		err = h.events.TurnStarted(ctx, gameID, events.TurnStartedEvent{
			Type:              "turnStarted",
			CurrentPlayerID:   playerID,
			Turn:              turn,
			TurnTimeRemaining: int(turnTimeLimit / time.Second),
		})
		if err != nil {
			return fmt.Errorf("send turn started event: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("turn start transaction: %w", err)
	}

	ticker := time.NewTicker(time.Millisecond * 500)

loop:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case t := <-ticker.C:
			currentTurn, err := h.repo.GetTurn(ctx, gameID)
			if err != nil {
				return fmt.Errorf("get current turn: %w", err)
			}
			remaining := currentTurn.EndAt.Sub(t)
			remainingSeconds := int(math.Round(remaining.Seconds()))
			if remainingSeconds == 0 {
				ticker.Stop()
				break loop
			}
			err = h.events.TurnTimeRemainingChanged(ctx, gameID, events.TurnTimeRemainingChangedEvent{
				Type:             "turnTimeRemaining",
				CurrentPlayerID:  playerID,
				RemainingSeconds: remainingSeconds,
			})
			if err != nil {
				return fmt.Errorf("send turn time remaining event: %w", err)
			}
		}
	}

	var nextPlayerID, nextTurn *int
	if turn+1 <= turnsCount {
		p := playerID ^ 1 // Toggle player ID (0 -> 1, 1 -> 0)
		nextPlayerID = &p
		t := turn + 1
		nextTurn = &t
	}

	err = h.events.TurnEnded(ctx, gameID, events.TurnEndedEvent{
		Type:         "turnEnded",
		NextPlayerID: nextPlayerID,
		NextTurn:     nextTurn,
	})
	if err != nil {
		return fmt.Errorf("send turn ended event: %w", err)
	}

	return nil
}
