package events

import (
	"context"

	"github.com/google/uuid"
)

type SubmissionSucceededEvent struct {
	Type       string `json:"type"`
	PlayerID   int    `json:"playerId"`
	Expression string `json:"expression"`
	Score      int    `json:"score"`
}

type Event interface {
	SubmissionSucceeded(ctx context.Context, gameID uuid.UUID, event SubmissionSucceededEvent) error
}
