package domain

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID        uuid.UUID
	Status    GameStatus
	StartedAt *time.Time
	EndedAt   *time.Time
	CreatedAt time.Time
}

type GameStatus string

const (
	GameStatusWaiting  GameStatus = "waiting"
	GameStatusStarted  GameStatus = "running"
	GameStatusFinished GameStatus = "finished"
	GameStatusCanceled GameStatus = "canceled"
)
