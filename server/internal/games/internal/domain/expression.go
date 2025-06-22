package domain

import (
	"time"

	"github.com/google/uuid"
)

type Expression struct {
	ID          uuid.UUID
	GameID      uuid.UUID
	PlayerID    int
	Expression  string
	Value       string
	Points      int
	Success     bool
	SubmittedAt time.Time
}
