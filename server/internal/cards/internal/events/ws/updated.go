package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/cards/internal/events"
	"github.com/traP-jp/h25s_15/internal/core/corews"
)

func (e *Event) CardUpdated(ctx context.Context, gameID uuid.UUID, event events.CardUpdatedEvent) error {
	var eventJSON bytes.Buffer
	err := json.NewEncoder(&eventJSON).Encode(event)
	if err != nil {
		return fmt.Errorf("encode event: %w", err)
	}

	err = e.m.BroadcastFilter(eventJSON.Bytes(), func(s *melody.Session) bool {
		if sessionGameIDI, ok := s.Get(corews.SessionKeyGameID); ok {
			if sessionGameID, ok := sessionGameIDI.(uuid.UUID); ok {
				return sessionGameID == gameID
			}
		}
		return false
	})
	if err != nil {
		return fmt.Errorf("broadcast event: %w", err)
	}

	return nil
}
