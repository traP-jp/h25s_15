package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/cards/internal/domain"
	"github.com/traP-jp/h25s_15/internal/games"
)

type EventCard struct {
	ID    uuid.UUID `json:"id"`
	Type  string    `json:"type"`
	Value int       `json:"value"`
}

type CardUpdatedEvent struct {
	Type              string      `json:"type"`
	FieldCards        []EventCard `json:"fieldCards"`
	Player0           []EventCard `json:"player0"`
	Player0HandsLimit int         `json:"player0HandsLimit"`
	Player1           []EventCard `json:"player1"`
	Player1HandsLimit int         `json:"player1HandsLimit"`
}

func (e *Event) CardUpdated(ctx context.Context, gameID uuid.UUID) error {
	cards, err := e.repo.GetActiveCards(ctx, gameID)
	if err != nil {
		return fmt.Errorf("get active cards: %w", err)
	}

	handLimits, err := e.repo.GetGameHandLimit(ctx, gameID)
	if err != nil {
		return fmt.Errorf("get game hand limit: %w", err)
	}

	fieldCards := make([]EventCard, 0, len(cards))
	player0Cards := make([]EventCard, 0, len(cards))
	player1Cards := make([]EventCard, 0, len(cards))

	for _, card := range cards {
		eventCard := EventCard{
			ID:    card.ID,
			Type:  string(card.Type),
			Value: card.Value,
		}

		switch card.Location {
		case domain.CardLocationField:
			fieldCards = append(fieldCards, eventCard)
		case domain.CardLocationHand:
			switch card.OwnerPlayerID {
			case 0:
				player0Cards = append(player0Cards, eventCard)
			case 1:
				player1Cards = append(player1Cards, eventCard)
			default:
				return fmt.Errorf("unexpected owner player ID: %d", card.OwnerPlayerID)
			}
		default:
			return fmt.Errorf("unexpected card location: %s", card.Location)
		}
	}

	event := CardUpdatedEvent{
		Type:              "cardUpdated",
		FieldCards:        fieldCards,
		Player0:           player0Cards,
		Player0HandsLimit: handLimits[0],
		Player1:           player1Cards,
		Player1HandsLimit: handLimits[1],
	}

	var eventJSON bytes.Buffer
	err = json.NewEncoder(&eventJSON).Encode(event)
	if err != nil {
		return fmt.Errorf("encode event: %w", err)
	}

	err = e.m.BroadcastFilter(eventJSON.Bytes(), func(s *melody.Session) bool {
		if sessionGameIDI, ok := s.Get(games.GameIDSessionKey); ok {
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
