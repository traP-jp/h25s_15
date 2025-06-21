package db

import "github.com/google/uuid"

type HandLimit struct {
	GameID         uuid.UUID `db:"game_id"`
	PlayerID       int       `db:"player_id"`
	HandCardsLimit int       `db:"hand_cards"`
}
