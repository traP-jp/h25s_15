package events

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ScoreUpdatedEvent struct {
	Type    string `json:"type"`
	Player0 int    `json:"player0"`
	Player1 int    `json:"player1"`
}

type Event interface {
	HandleRequestWithKeys(res http.ResponseWriter, req *http.Request, keys map[string]any)
	ScoreUpdated(ctx context.Context, gameID uuid.UUID, event ScoreUpdatedEvent) error
}
