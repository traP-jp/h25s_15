package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (e *Event) GameReady(ctx context.Context, gameID uuid.UUID, event events.GameReadyEvent, playerNames [2]string) error {
	sessions, err := e.m.Sessions()
	if err != nil {
		return fmt.Errorf("get sessions: %w", err)
	}

	filteredSessions := make([]*melody.Session, 0, len(sessions))
	for _, sess := range sessions {
		if corews.FilterGameID(gameID)(sess) {
			filteredSessions = append(filteredSessions, sess)
		}
	}

	event0 := event
	event0.PlayerID = 0
	event0JSON, err := corews.JSON(event0)
	if err != nil {
		return fmt.Errorf("marshal event0: %w", err)
	}
	event1 := event
	event1.PlayerID = 1
	event1JSON, err := corews.JSON(event1)
	if err != nil {
		return fmt.Errorf("marshal event1: %w", err)
	}

	sess0 := make([]*melody.Session, 0, len(filteredSessions))
	sess1 := make([]*melody.Session, 0, len(filteredSessions))
	for _, sess := range filteredSessions {
		if corews.FilterUserName(playerNames[0])(sess) {
			sess0 = append(sess0, sess)
		} else if corews.FilterUserName(playerNames[1])(sess) {
			sess1 = append(sess1, sess)
		}
	}

	if len(sess0) > 0 {
		err := e.m.BroadcastMultiple(event0JSON, sess0)
		if err != nil {
			return fmt.Errorf("broadcast event0: %w", err)
		}
	}
	if len(sess1) > 0 {
		err := e.m.BroadcastMultiple(event1JSON, sess1)
		if err != nil {
			return fmt.Errorf("broadcast event1: %w", err)
		}
	}

	return nil
}
