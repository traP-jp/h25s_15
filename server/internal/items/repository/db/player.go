package db

import "github.com/google/uuid"

type Player struct {
	GameID   uuid.UUID `db:"game_id"`
	PlayerID int       `db:"player_id"`
	UserName string    `db:"user_name"`
	Score    int       `db:"score"`
}
