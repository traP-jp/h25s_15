package domain

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	GameID   uuid.UUID
	PlayerID int
	UserName string
	Score    int
}

type WaitingPlayer struct {
	UserName  string
	CreatedAt time.Time
}
