package domain

import "github.com/google/uuid"

type GamePlayer struct {
	GameID   uuid.UUID
	PlayerID int
	UserName string
	Score    int
}
