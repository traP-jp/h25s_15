package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID        uuid.UUID    `db:"id"`
	Status    string       `db:"status"`
	StartedAt sql.NullTime `db:"started_at"`
	EndedAt   sql.NullTime `db:"ended_at"`
	CreatedAt time.Time    `db:"created_at"`
}
