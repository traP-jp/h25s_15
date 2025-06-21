package domain

import (
	"time"

	"github.com/google/uuid"
)

type Turn struct {
	GameID     uuid.UUID
	PlayerID   int
	TurnNumber int
	StartAt    time.Time
	EndAt      time.Time
}
