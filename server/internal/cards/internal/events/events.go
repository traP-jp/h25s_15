package events

import (
	"context"

	"github.com/google/uuid"
)

type Event interface {
	CardUpdated(ctx context.Context, gameID uuid.UUID, event CardUpdatedEvent) error
}

type EventCard struct {
	ID    uuid.UUID `json:"id"`
	Type  string    `json:"type"`
	Value string    `json:"value"`
}

type CardUpdatedEvent struct {
	Type              string      `json:"type"`
	FieldCards        []EventCard `json:"fieldCards"`
	Player0           []EventCard `json:"player0"`
	Player0HandsLimit int         `json:"player0HandsLimit"`
	Player1           []EventCard `json:"player1"`
	Player1HandsLimit int         `json:"player1HandsLimit"`
}
