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

type GameMatchedEvent struct {
	GameID   uuid.UUID `json:"gameId"`
	PlayerID int       `json:"playerId"`
}

type Event interface {
	HandleRequestWithKeys(res http.ResponseWriter, req *http.Request, keys map[string]any)
	ScoreUpdated(ctx context.Context, gameID uuid.UUID, event ScoreUpdatedEvent) error
	GameMatched(ctx context.Context, userNames [2]string, gameID uuid.UUID) error
	GetConnectedWaitingUsers(ctx context.Context) ([]string, error)
}
