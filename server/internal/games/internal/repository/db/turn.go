package db

import (
	"time"

	"github.com/google/uuid"
)

type Turn struct {
	GameID     uuid.UUID `db:"game_id"`
	PlayerID   int       `db:"player_id"`
	TurnNumber int       `db:"turn_number"`
	StartAt    time.Time `db:"start_at"`
	EndAt      time.Time `db:"end_at"`
}
