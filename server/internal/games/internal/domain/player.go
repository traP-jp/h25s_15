package domain

import "github.com/google/uuid"

type Player struct {
	GameID   uuid.UUID
	PlayerID int
	UserName string
	Score    int
}
