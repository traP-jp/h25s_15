package events

import (
	"context"
	"net/http"
	"time"

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

type EventCard struct {
	ID    uuid.UUID `json:"id"`    // Card ID
	Type  string    `json:"type"`  // Card type (e.g., "operator", "operand", "item")
	Value string    `json:"value"` // Card value (e.g., "1", "2", "3", etc.)
}

type GameReadyEvent struct {
	Type              string      `json:"type"`
	FieldCards        []EventCard `json:"fieldCards"`        // Card IDs in the field
	PlayerID          int         `json:"playerId"`          // Player ID of the user who is ready
	Player0           []EventCard `json:"player0"`           // Cards in player 0's hand
	Player0HandsLimit int         `json:"player0HandsLimit"` // Hand limit for player 0
	Player1           []EventCard `json:"player1"`           // Cards in player 1's hand
	Player1HandsLimit int         `json:"player1HandsLimit"` // Hand limit for player 1
	CurrentPlayerID   int         `json:"currentPlayerId"`   // ID of the player whose turn it is
	Player0Score      int         `json:"player0Score"`      // Score of player 0
	Player1Score      int         `json:"player1Score"`      // Score of player 1
	StartTime         time.Time   `json:"startTime"`         // Start time of the game
}

type GameStartedEvent struct {
	Type            string `json:"type"`
	CurrentPlayerID int    `json:"currentPlayerId"` // ID of the player whose turn it is
	Turn            int    `json:"turn"`            // Current turn number
	TurnTotal       int    `json:"turnTotal"`       // Total number of turns in the game
}

type TurnStartedEvent struct {
	Type              string `json:"type"`
	CurrentPlayerID   int    `json:"currentPlayerId"`   // ID of the player whose turn it is
	Turn              int    `json:"turn"`              // Current turn number
	TurnTimeRemaining int    `json:"turnTimeRemaining"` // Remaining time for the turn in seconds
}

type TurnEndedEvent struct {
	Type         string `json:"type"`
	NextPlayerID *int   `json:"nextPlayerId"` // ID of the player whose turn it is
	NextTurn     *int   `json:"nextTurn"`     // Next turn number
}

type TurnTimeRemainingChangedEvent struct {
	Type             string `json:"type"`
	CurrentPlayerID  int    `json:"currentPlayerId"`  // ID of the player whose turn it is
	RemainingSeconds int    `json:"remainingSeconds"` // Remaining time for the turn in seconds
}

type GameEndedEvent struct {
	Type    string `json:"type"`
	Player0 int    `json:"player0"` // Score of player 0
	Player1 int    `json:"player1"` // Score of player 1
}

type Event interface {
	HandleRequestWithKeys(res http.ResponseWriter, req *http.Request, keys map[string]any)
	GetConnectedWaitingUsers(ctx context.Context) ([]string, error)
	// GetGameConnectedUsers returns a list of user names connected to the game with the given gameID.
	GetGameConnectedUsers(ctx context.Context, gameID uuid.UUID) ([]string, error)

	ScoreUpdated(ctx context.Context, gameID uuid.UUID, event ScoreUpdatedEvent) error
	GameMatched(ctx context.Context, userNames [2]string, gameID uuid.UUID) error
	// 内部でPlayerIDをよしなに付けてイベントを送る
	GameReady(ctx context.Context, gameID uuid.UUID, event GameReadyEvent, playerNames [2]string) error
	GameStarted(ctx context.Context, gameID uuid.UUID, event GameStartedEvent) error
	TurnStarted(ctx context.Context, gameID uuid.UUID, event TurnStartedEvent) error
	TurnEnded(ctx context.Context, gameID uuid.UUID, event TurnEndedEvent) error
	TurnTimeRemainingChanged(ctx context.Context, gameID uuid.UUID, event TurnTimeRemainingChangedEvent) error
	GameEnded(ctx context.Context, gameID uuid.UUID, event GameEndedEvent) error
}
