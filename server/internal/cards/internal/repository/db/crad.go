package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Card struct {
	ID            uuid.UUID     `db:"id"`
	GameID        uuid.UUID     `db:"game_id"`
	Type          string        `db:"type"`
	Value         string        `db:"value"`
	OwnerPlayerID sql.NullInt16 `db:"owner_player_id"`
	Location      string        `db:"location"`
}