package db

import (
	"time"

	"github.com/google/uuid"
)

type Expression struct {
	ID          uuid.UUID `db:"id"`
	GameID      uuid.UUID `db:"game_id"`
	PlayerID    int       `db:"player_id"`
	Expression  string    `db:"expression"`
	Value       string    `db:"value"`
	Points      int       `db:"points"`
	Success     bool      `db:"success"`
	SubmittedAt time.Time `db:"submitted_at"` // Use string for compatibility with database format
}
